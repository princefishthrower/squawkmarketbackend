package jobs

import (
	"log"
	"squawkmarketbackend/scraper"
	scraperTypes "squawkmarketbackend/scraper/types"
	"time"

	"github.com/philippseith/signalr"
)

func StartHeadlineScrapeJob(server signalr.Server) {
	log.Println("Starting Headline Scrape Job")
	// start in a goroutine so it doesn't block
	go func() {
		scraper.ScrapeForConfigs(server, scraperTypes.HeadlineScrapingConfigs, 10*time.Second)
	}()
}
