package jobs

import (
	"log"
	"squawkmarketbackend/db"
	"squawkmarketbackend/googletexttospeech"
	"squawkmarketbackend/hub"
	"time"

	"github.com/philippseith/signalr"
	"github.com/robfig/cron/v3"
)

func StartMarketOpenJob(server signalr.Server, est *time.Location) {
	feedName := "market-wide"
	c := cron.New(cron.WithLocation(est))

	// run postmarket cron every weekday at 5:00pm EST
	c.AddFunc("9 30 * * 1-5", func() {

		// generate premarket message
		marketOpenMessage := "The market is now open."

		// convert to MP3
		mp3Data := googletexttospeech.TextToSpeech(marketOpenMessage)

		// insert into database
		err := db.InsertSquawkIfNotExists("", "", feedName, marketOpenMessage, mp3Data)
		if err != nil {
			log.Println("Error inserting squawk into database:", err)
			return
		}

		squawk, err := db.GetLatestSquawk()
		if err != nil {
			log.Println("Error getting latest squawk from database:", err)
			return
		}
		// ship the latest squawk over the WebSocket
		hub.BroadcastSquawk(server, feedName, squawk)
	})
	c.Start()
	log.Println("Started Market Open Cron Job")
}
