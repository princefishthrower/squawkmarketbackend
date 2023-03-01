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

	// run postmarket cron every weekday at 16:02pm EST
	c.AddFunc("2 16 * * 1-5", func() {

		// generate premarket message
		postmarketMessage, err := generators.GeneratePostmarketMessage()
		if err != nil {
			log.Println("Error generating premarket message: ", err)
			return
		}

		// convert to MP3
		mp3Data := googletexttospeech.TextToSpeech(*postmarketMessage)

		// insert into database
		err = db.InsertSquawk("", "", feedName, *postmarketMessage, mp3Data)
		if err != nil {
			log.Println("Error inserting squawk into database:", err)
			return
		}

		squawk, err := db.GetLatestSquawkByFeed("market-wide")
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
