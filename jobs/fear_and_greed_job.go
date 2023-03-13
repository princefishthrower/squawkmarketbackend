package jobs

import (
	"log"
	"squawkmarketbackend/fear_and_greed"
	"squawkmarketbackend/scraper"

	"time"

	"github.com/philippseith/signalr"
	"github.com/robfig/cron/v3"
)

func FearAndGreedJob(server signalr.Server, est *time.Location) {
	c := cron.New(cron.WithLocation(est))

	// every minute, get the fear and greed index squawk
	c.AddFunc("* 8-17 * * 1-5", func() {
		squawk, err := fear_and_greed.GetFearAndGreedSquawk()
		if err != nil {
			return
		}

		// send the squawk to the hub
		scraper.GenerateAndStoreFeedItemIfNotExists(squawk, "", "market-wide", 0, server)
	})
	c.Start()
	log.Println("Started Fear and Green Cron Job")
}
