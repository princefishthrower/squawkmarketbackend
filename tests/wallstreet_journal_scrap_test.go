package tests

import (
	"squawkmarketbackend/scraper"
	scraperTypes "squawkmarketbackend/scraper/types"
	"testing"
)

func TestWallstreetJournalScrape(t *testing.T) {
	squawk, err := scraper.ScrapeForConfigItem(scraperTypes.WallStreetJournalNewsConfig)
	if err != nil {
		t.Errorf("Error scraping for Wall Street Journal: %v", err)
	}

	if squawk == nil {
		t.Errorf("Squawk is nil")
		return
	}

	if *squawk == "" {
		t.Errorf("Squawk is empty")
		return
	}

	t.Logf("TEST WALL STREET JOURNAL\n\n%v", squawk)
}
