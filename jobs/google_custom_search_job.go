package jobs

import (
	"log"
	"squawkmarketbackend/db"
	"squawkmarketbackend/googlecustomsearch"
	"squawkmarketbackend/googletexttospeech"
	"squawkmarketbackend/hub"

	"github.com/philippseith/signalr"
	"github.com/robfig/cron/v3"
)

func StartGoogleCustomSearchJob(server signalr.Server) {
	feedName := "market-wide"

	c := cron.New(cron.WithSeconds())

	// run custom search every 10 seconds
	c.AddFunc("*/10 * * * * *", func() {

		// run custom google search
		squawk, err := googlecustomsearch.CustomSearch("financial breaking news")
		if err != nil {
			log.Println("Error running custom search:", err)
			return
		}

		squawkExists, err := db.DoesSquawkExistAccordingToFeedCriterion(squawk.Squawk, "", feedName, 0.75)
		if err != nil {
			log.Println("Error checking if squawk exists:", err)
			return
		}
		if squawkExists {
			log.Println("Squawk already exists in database, skipping")
			return
		}

		// convert to MP3
		mp3Data := googletexttospeech.TextToSpeech(squawk.Squawk)

		// insert into database
		err = db.InsertSquawk("", "", feedName, squawk.Squawk, mp3Data)
		if err != nil {
			log.Println("Error inserting squawk into database:", err)
			return
		}

		squawkObj, err := db.GetLatestSquawkByFeed("market-wide")
		if err != nil {
			log.Println("Error getting latest squawk from database:", err)
			return
		}

		// ship the latest squawk over the WebSocket
		hub.BroadcastSquawk(server, feedName, squawkObj)
	})
	c.Start()
	log.Println("Started Custom Search Cron Job")
}
