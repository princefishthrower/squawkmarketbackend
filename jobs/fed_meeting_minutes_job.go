package jobs

import (
	"log"
	"squawkmarketbackend/db"
	"squawkmarketbackend/generators"
	"squawkmarketbackend/googletexttospeech"
	"squawkmarketbackend/hub"

	"github.com/philippseith/signalr"
	"github.com/robfig/cron/v3"
)

func StartFedMeetingMinutesJob(server signalr.Server) {
	feedName := "market-wide"
	c := cron.New()
	// run fed minutes cron every minute
	c.AddFunc("8 * * * *", func() {
		fedMeetingMinutesSummary, err := generators.GenerateFedMeetingMinutesSummary()
		if err != nil {
			log.Println("Error generating premarket message: ", err)
			return
		}

		// convert to MP3
		mp3Data := googletexttospeech.TextToSpeech(*fedMeetingMinutesSummary)

		// insert into database
		err = db.InsertSquawk("", "", feedName, *fedMeetingMinutesSummary, mp3Data)
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
	log.Println("Started Fed Meeting Minutes Cron Job")
}
