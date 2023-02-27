package tests

import (
	"squawkmarketbackend/generators"
	"testing"
)

func TestGenerateFedMeetingMinutes(t *testing.T) {
	fedMeetingMinutesSummary, err := generators.GenerateFedMeetingMinutesSummary()
	if err != nil {
		t.Errorf("Error generating Fed Meeting Minutes Summary: %s", err)
	}
	if fedMeetingMinutesSummary == nil {
		t.Errorf("Fed Meeting Minutes Summary is nil")
	}
	if *fedMeetingMinutesSummary == "" {
		t.Errorf("Fed Meeting Minutes Summary is empty")
	}

	t.Logf("IN TEST: BEGIN FED MEETING MINUTES SUMMARY:\n\n%s", *fedMeetingMinutesSummary)
}
