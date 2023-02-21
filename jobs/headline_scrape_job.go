package jobs

import (
	"log"
	"squawkmarketbackend/scraper"

	"github.com/philippseith/signalr"
)

func StartFeedItemScrapeJob(server signalr.Server) {
	log.Println("Starting Feed Item Scrape Job")
	// start in a goroutine so it doesn't block
	go func() {
		scraper.ParseFeedItems(server)
	}()
}
