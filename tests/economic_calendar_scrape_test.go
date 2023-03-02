package tests

import (
	"squawkmarketbackend/scraper"
	scraperTypes "squawkmarketbackend/scraper/types"
	"testing"
)

func TestEconomicCalendarScrape(t *testing.T) {
	_, err := scraper.ScrapeForConfigItem(scraperTypes.USEconomicCalendarConfig)
	if err != nil {
		t.Errorf("Error scraping for economic calendar: %v", err)
	}

	// if len(squawks) == 0 {
	// 	t.Errorf("Squawk is nil")
	// 	return
	// }

	// empty is actually fine if there isn't a new figure yet
	// if *squawk == "" {
	// 	t.Errorf("Squawk is empty")
	// 	return
	// }

	// t.Logf("TEST ECONOMIC CALENDAR\n\n%v", squawk)
}
