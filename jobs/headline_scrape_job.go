package jobs

import (
	"squawkmarketbackend/headlines"

	"github.com/philippseith/signalr"
	"github.com/robfig/cron/v3"
)

func StartHeadlineScrapeJob(server signalr.Server) {
	// Run every 10 seconds
	c := cron.New(cron.WithSeconds())
	c.AddFunc("*/10 * * * * *", func() {
		//  for each config in the slice of configs
		headlines.ParseHeadlines(server)
	})

	c.Start()
}
