package tests

import (
	browser_scraper "squawkmarketbackend/browser_scraper/handlers"

	"testing"
)

func TestScrapeNasdaqTradingHalts(t *testing.T) {
	squawk, err := browser_scraper.GetLatestNasdaqTradingHalt()
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if *squawk == "" {
		t.Errorf("Error: squawk is empty")
	}

	// log the squawk
	t.Logf("Squawk: %v", squawk)
}
