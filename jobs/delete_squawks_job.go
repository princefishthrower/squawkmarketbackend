package jobs

import (
	"log"
	"squawkmarketbackend/db"

	"time"

	"github.com/philippseith/signalr"
	"github.com/robfig/cron/v3"
)

func DeleteSquawksJob(server signalr.Server, est *time.Location) {
	c := cron.New(cron.WithLocation(est))

	// every day at 8 PM EST, delete all squawks from the database
	c.AddFunc("0 20 * * *", func() {
		err := db.DeleteAllSquawks()
		if err != nil {
			log.Println("Error deleting all squawks from database:", err)
			return
		}
	})
	c.Start()
	log.Println("Started Delete Squawks Cron Job")
}
