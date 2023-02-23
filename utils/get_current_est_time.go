package utils

import "time"

func GetCurrentESTTime() string {
	// Set the location to Eastern Standard Time
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		panic(err)
	}

	// Get the current time in EST
	now := time.Now().In(loc)

	// Format the time as a string
	timeString := now.Format("3:04 PM")

	return timeString
}
