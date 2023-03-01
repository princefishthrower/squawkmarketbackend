package tests

import (
	"log"
	"squawkmarketbackend/generators"
	"testing"

	"github.com/joho/godotenv"
)

func TestGenerateFedMeetingMinutes(t *testing.T) {

	// need env for open AI key
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fedMeetingMinutesSummary, err := generators.GenerateFedMeetingMinutesSummary()
	if err != nil {
		t.Errorf("Error generating Fed Meeting Minutes Summary: %s", err)
		return
	}
	if fedMeetingMinutesSummary == nil {
		t.Errorf("Fed Meeting Minutes Summary is nil")
		return
	}
	if *fedMeetingMinutesSummary == "" {
		t.Errorf("Fed Meeting Minutes Summary is empty")
		return
	}

	t.Logf("IN TEST: BEGIN FED MEETING MINUTES SUMMARY:\n\n%s", *fedMeetingMinutesSummary)
}
