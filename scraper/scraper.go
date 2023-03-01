package scraper

import (
	"log"
	"squawkmarketbackend/db"
	"squawkmarketbackend/googletexttospeech"
	"squawkmarketbackend/hub"
	scraperTypes "squawkmarketbackend/scraper/types"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/philippseith/signalr"
)

func ScrapeForConfigItems(server signalr.Server) {
	for _, config := range scraperTypes.ScrapingConfigs {
		// get the squawks from the config
		squawks, err := ScrapeForConfigItem(config)
		if err != nil {
			log.Println("Error getting squawk:", err)
			return
		}

		// log out squawks length
		log.Println("Squawks found and parsed, from: ", config.Url, ": ", len(squawks))

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
		time.Sleep(5 * time.Second)
	}

	// and then start all over again
	ScrapeForConfigItems(server)
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
	squawkExists, err := db.DoesSquawkExistAccordingToFeedCriterion(squawk, symbols, feedName, insertThreshold)
	if err != nil {
		log.Println("Error checking if squawk exists:", err)
		return
	}
	if squawkExists {
		log.Println("Squawk already exists in database, skipping")
		return
	}

	// generate MP3 data
	// mp3Data := elevenlabs.TextToSpeech(squawk)
	mp3Data := googletexttospeech.TextToSpeech(squawk)

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
