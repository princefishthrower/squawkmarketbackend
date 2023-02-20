package headlines

import (
	"log"
	"squawkmarketbackend/db"
	"squawkmarketbackend/elevenlabs"
	headlinesTypes "squawkmarketbackend/headlines/types"
	"squawkmarketbackend/hub"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/philippseith/signalr"
)

func ParseHeadlines(server signalr.Server) {
	for _, config := range headlinesTypes.HeadlineConfigs {
		// get the headline
		headline, err := ParseHeadline(config.Url, config.Selector)
		if err != nil {
			log.Println("Error getting headline:", err)
			return
		}
		// ship and store headline
		GenerateAndStoreHeadlineIfNotExists(*headline, server)
	}
}

func ParseHeadline(url string, selector string) (*string, error) {

	c := colly.NewCollector(
		colly.AllowedDomains("finviz.com", "finance.yahoo.com", "marketwatch.com", "reuters.com", "wsj.com"),
	)

	headline := ""
	c.OnHTML(selector, func(e *colly.HTMLElement) {
		// only get the first headline
		if headline == "" {
			headline = e.Text
			// remove all leading and trailing whitespace
			headline = strings.TrimSpace(headline)

			// // remove all newlines and tabs
			headline = strings.ReplaceAll(headline, "\n", "")
			headline = strings.ReplaceAll(headline, "\t", "")

			// // regex remove any 00:00AM / 00:00PM times (military time)
			headline = strings.ReplaceAll(headline, "[0-9][0-9]:[0-9][0-9][AP]M", "")

			log.Println("Headline found and parsed, from: ", url, ": ", headline)
			return
		}
	})

	c.Visit(url)

	return &headline, nil
}

func GenerateAndStoreHeadlineIfNotExists(headline string, server signalr.Server) {
	// check if headline is already in database
	headlineExists, err := db.DoesHeadlineExist(headline)
	if err != nil {
		log.Println("Error checking if headline exists:", err)
		return
	}
	if headlineExists {
		log.Println("Headline already exists in database:", headline)
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
