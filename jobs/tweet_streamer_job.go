package jobs

import (
	"log"
	"squawkmarketbackend/twitter"
	"time"

	"github.com/philippseith/signalr"
)

func StartTweetStreamerJob(server signalr.Server, est *time.Location) {
	log.Println("Starting Tweet Streamer Job")
	go func() {
		// get tweets from financialjuice
		twitter.StreamTweetsOfUser("financialjuice")

		// also from FirstSquawk
		twitter.StreamTweetsOfUser("FirstSquawk")
	}()
}
