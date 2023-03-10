package jobs

import (
	"log"
	"squawkmarketbackend/scraper"
	scraperTypes "squawkmarketbackend/scraper/types"
	"time"

	"github.com/philippseith/signalr"
)

func StartSectorJob(server signalr.Server, est *time.Location) {
	log.Println("Starting Finviz Sector Job")
	// at minute intervals get change in sectors
	go func() {
		scraper.ScrapeForConfigs(server, []scraperTypes.ScrapingConfig{scraperTypes.FinvizSectorConfig}, 60*time.Second)
	}()
}
