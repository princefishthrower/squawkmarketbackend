package jobs

import (
	"log"
	"squawkmarketbackend/db"
	"squawkmarketbackend/generators"
	"squawkmarketbackend/googletexttospeech"
	"squawkmarketbackend/hub"
	"time"

	"github.com/philippseith/signalr"
	"github.com/robfig/cron/v3"
)

func StartPostmarketJob(server signalr.Server, est *time.Location) {
	feedName := "market-wide"
	c := cron.New(cron.WithLocation(est))

	// run postmarket cron every weekday at 5:00pm EST
	c.AddFunc("0 17 * * 1-5", func() {

		// generate premarket message
		postmarketMessage, err := generators.GeneratePostmarketMessage()
		if err != nil {
			log.Println("Error generating premarket message: ", err)
			return
		}

		// convert to MP3
		mp3Data := googletexttospeech.TextToSpeech(*postmarketMessage)

		// insert into database
		err = db.InsertSquawkIfNotExists("", "", feedName, *postmarketMessage, mp3Data)

		squawk, err := db.GetLatestSquawk()
		if err != nil {
			log.Println("Error getting latest squawk from database:", err)
			return
		}
		// ship the latest squawk over the WebSocket
		hub.BroadcastSquawk(server, feedName, squawk)
	})
	c.Start()
	log.Println("Started Premarket Cron Job")
}
