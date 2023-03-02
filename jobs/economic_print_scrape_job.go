package jobs

import (
	"log"
	"squawkmarketbackend/scraper"
	scraperTypes "squawkmarketbackend/scraper/types"
	"time"

	"github.com/philippseith/signalr"
)

func EconomicPrintScrapeJob(server signalr.Server) {
	log.Println("Starting Economic Prints Scrape Job")
	// start in a goroutine so it doesn't block
	go func() {
		scraper.ScrapeForConfigs(server, scraperTypes.EconomicPrintScrapingConfigs, 5*time.Second)
	}()
}
