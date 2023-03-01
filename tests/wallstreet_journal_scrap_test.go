package tests

import (
	"squawkmarketbackend/scraper"
	scraperTypes "squawkmarketbackend/scraper/types"
	"testing"
)

func TestWallstreetJournalScrape(t *testing.T) {
	squawks, err := scraper.ScrapeForConfigItem(scraperTypes.WallStreetJournalNewsConfig)
	if err != nil {
		t.Errorf("Error scraping for Wall Street Journal: %v", err)
	}

	if len(squawks) == 0 {
		t.Errorf("Squawks is empty")
		return
	}

	if squawks[0] == "" {
		t.Errorf("squawks[0] is empty")
		return
	}

	t.Logf("TEST WALL STREET JOURNAL\n\n%v", squawks[0])
}
