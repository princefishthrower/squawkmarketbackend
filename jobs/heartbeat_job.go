package jobs

import (
	"log"
	"squawkmarketbackend/http_helper"

	"github.com/robfig/cron/v3"
)

func StartHeartBeatJob() {
	c := cron.New()

	// run every minute
	c.AddFunc("* * * * * ", func() {
		_, err := http_helper.MakeHTTPRequest("https://betteruptime.com/api/v1/heartbeat/e2Fz4BBhGyeWiGgPuwqF7d99", "POST", nil, nil, nil)
		if err != nil {
			log.Println(err)
		}
	})
	log.Println("Started cron job for heartbeat")
	c.Start()
}
