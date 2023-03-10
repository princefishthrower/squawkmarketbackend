package jobs

import (
	"log"
	"squawkmarketbackend/scraper"
	scraperTypes "squawkmarketbackend/scraper/types"
	"time"

	"github.com/philippseith/signalr"
)

func StartFinvizScrapeJob(server signalr.Server) {
	log.Println("Starting Finviz Technical Indicators Scrape Job")
	// start in a goroutine so it doesn't block
	go func() {
		scraper.ScrapeForConfigs(server, scraperTypes.FinvizTechnicalIndicatorsScrapingConfigs, 10*time.Second)
	}()
}
