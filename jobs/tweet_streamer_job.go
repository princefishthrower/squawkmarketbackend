package jobs

import (
	"log"
	"squawkmarketbackend/twitter"
	"time"

	"github.com/philippseith/signalr"
)

func StartTweetStreamerJob(server signalr.Server, est *time.Location) {
	log.Println("Starting Tweet Streamer Job")
	// at minute intervals get change in sectors
	go func() {
		twitter.StreamTweetsOfUser("financialjuice")
	}()
}
