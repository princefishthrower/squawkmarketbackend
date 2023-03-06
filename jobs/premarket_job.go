package jobs

import (
	"log"
	"squawkmarketbackend/amazontexttospeech"
	"squawkmarketbackend/db"
	"squawkmarketbackend/generators"
	"squawkmarketbackend/hub"
	"time"

	"github.com/philippseith/signalr"
	"github.com/robfig/cron/v3"
)

func StartPremarketJob(server signalr.Server, est *time.Location) {
	feedName := "market-wide"

	c := cron.New(cron.WithLocation(est))

	// run cron every weekday at 9:25am
	c.AddFunc("25 9 * * 1-5", func() {

		// generate premarket message
		premarketMessage, err := generators.GeneratePremarketMessage()
		if err != nil {
			log.Println("Error generating premarket message: ", err)
			return
		}

		// convert to MP3
		mp3Data, err := amazontexttospeech.TextToSpeech(*premarketMessage)
		if err != nil {
			log.Println("Error converting text to speech:", err)
			return
		}

		// insert into database
		err = db.InsertSquawk("", "", feedName, *premarketMessage, mp3Data)
		if err != nil {
			log.Println("Error inserting squawk into database:", err)
			return
		}

		squawk, err := db.GetLatestSquawkByFeed(feedName)
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
