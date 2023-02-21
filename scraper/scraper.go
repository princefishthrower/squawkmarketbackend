package scraper

import (
	"log"
	"squawkmarketbackend/db"
	"squawkmarketbackend/elevenlabs"
	"squawkmarketbackend/hub"
	scraperTypes "squawkmarketbackend/scraper/types"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/philippseith/signalr"
)

func ParseFeedItems(server signalr.Server) {
	for _, config := range scraperTypes.HeadlineConfigs {
		// get the headline
		headline, err := ParseFeedItem(config.Url, config.Selector, config.HandlerFunction)
		if err != nil {
			log.Println("Error getting headline:", err)
			return
		}
		log.Println("Headline found and parsed, from: ", config.Url, ": ", *headline)
		// ship and store headline
		GenerateAndStoreFeedItemIfNotExists(*headline, server)

		// wait 1 second before scraping the next headline
		time.Sleep(1 * time.Second)
	}

	// and then start all over again
	ParseFeedItems(server)
}

func ParseFeedItem(url string, selector string, onHtmlHandler func(*string, string) func(e *colly.HTMLElement)) (*string, error) {
	c := colly.NewCollector(
		colly.AllowedDomains("finviz.com", "finance.yahoo.com", "marketwatch.com", "reuters.com", "wsj.com"),
	)

	headline := ""
	c.OnHTML(selector, onHtmlHandler(&headline, url))
	c.Visit(url)

	return &headline, nil
}

func GenerateAndStoreFeedItemIfNotExists(headline string, server signalr.Server) {
	// check if headline is already in database
	headlineExists, err := db.DoesHeadlineExist(headline)
	if err != nil {
		log.Println("Error checking if headline exists:", err)
		return
	}
	if headlineExists {
		return
	}

	// generate MP3 data and send over WebSocket
	mp3Data := elevenlabs.GenerateMP3Data(headline)

	// add headline to database - will only add if the title is not already found in the database
	err = db.AddHeadline(headline, mp3Data)
	if err != nil {
		log.Println("Error adding headline to database:", err)
		return
	}

	latestEntry, err := db.GetLatestHeadline()
	if err != nil {
		log.Println("Error getting latest headline from database:", err)
		return
	}

	// ship the latest headline over the WebSocket
	hub.BroadcastHeadline(latestEntry, server)
}
