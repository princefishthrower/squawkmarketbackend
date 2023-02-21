package jobs

import (
	"log"
	"squawkmarketbackend/scraper"

	"github.com/philippseith/signalr"
)

func StartFeedItemScrapeJob(server signalr.Server) {
	log.Println("Starting Feed Item Scrape Job")
	scraper.ParseFeedItems(server)
}
