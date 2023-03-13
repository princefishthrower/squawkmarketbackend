package scraper

import (
	"log"
	"os"
	"squawkmarketbackend/amazontexttospeech"
	"squawkmarketbackend/db"
	"squawkmarketbackend/hub"
	scraperTypes "squawkmarketbackend/scraper/types"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/philippseith/signalr"
)

func ScrapeForConfigs(server signalr.Server, scrapingConfigs []scraperTypes.ScrapingConfig, duration time.Duration) {
	// if it is not between 8am and 5pm EST, don't scrape
	// if !IsItTimeToScrape() {
	// 	log.Println("It is not time to scrape, sleeping for 1 minute")
	// 	time.Sleep(1 * time.Minute)
	// 	ScrapeForConfigs(server, scrapingConfigs, duration)
	// 	return
	// }

	for _, config := range scrapingConfigs {
		// get the squawks from the config
		squawks, err := ScrapeForConfigItem(config)
		if err != nil {
			log.Println("Error getting squawk:", err)
			return
		}

		// loop at all squawks
		for _, squawk := range squawks {
			log.Println("Squawks found and parsed, from: ", config.Url, ": ", squawk)

			if squawk == "" {
				log.Println("Squawk is empty, skipping")
				continue
			}

			// ship and store squawk
			// TODO: add some sort of symbol service to determine relevant symbols
			GenerateAndStoreFeedItemIfNotExists(squawk, "", config.FeedName, config.InsertThreshold, server)
		}
		// wait 5 second before scraping the next squawk
		time.Sleep(duration)
	}

	// and then start all over again
	ScrapeForConfigs(server, scrapingConfigs, duration)
}

func IsItTimeToScrape() bool {

	if os.Getenv("ENVIRONMENT") != "production" {
		log.Println("Not in production, so we are always scraping")
		return true
	}

	est, err := time.LoadLocation("America/New_York")
	if err != nil {
		log.Println(err)
		return false
	}
	now := time.Now().In(est)
	if now.Hour() < 8 || now.Hour() > 17 {
		return false
	}
	return true
}

func ScrapeForConfigItem(config scraperTypes.ScrapingConfig) ([]string, error) {
	c := colly.NewCollector(
		colly.AllowedDomains(scraperTypes.AllowedDomains...),
		// useful for debugging
		// colly.Debugger(&debug.LogDebugger{}),
	)
	squawks := make([]string, 0)
	c.OnHTML(config.Selector, config.HandlerFunction(&squawks, config.Url))

	// Set error handler
	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.Visit(config.Url)
	return squawks, nil
}

func GenerateAndStoreFeedItemIfNotExists(squawk string, symbols string, feedName string, insertThreshold float64, server signalr.Server) {
	// check if squawk is already in database
	squawkExists, err := db.DoesSquawkAlreadyExistAccordingToFeedCriterion(squawk, symbols, feedName, insertThreshold)
	if err != nil {
		log.Println("Error checking if squawk exists:", err)
		return
	}
	if squawkExists {
		// log.Println("Squawk already exists in database, not generating mp3 or broadcasting")
		return
	}

	log.Println("Squawk does not exist in database, generating mp3 and broadcasting")

	// convert to MP3
	mp3Data, err := amazontexttospeech.TextToSpeech(squawk)
	if err != nil {
		log.Println("Error converting text to speech:", err)
		return
	}

	// add squawk to database - will only add if the title is not already found in the database
	err = db.InsertSquawk("", "", feedName, squawk, mp3Data)
	if err != nil {
		log.Println("Error adding squawk to database:", err)
		return
	}

	squawkObj, err := db.GetLatestSquawkByFeed(feedName)
	if err != nil {
		log.Println("Error getting latest squawk from database:", err)
		return
	}

	// ship the latest squawk over the WebSocket
	hub.BroadcastSquawk(server, feedName, squawkObj)
}
