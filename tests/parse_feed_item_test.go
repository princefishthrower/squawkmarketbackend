package tests

import (
	"fmt"
	"reflect"
	"runtime"
	"squawkmarketbackend/scraper"
	scraperTypes "squawkmarketbackend/scraper/types"

	"testing"
)

func TestParseFeedItem(t *testing.T) {
	// loop over all configs and ensure they are working
	for _, config := range scraperTypes.HeadlineConfigs {
		// use reflection to get the function name
		function := reflect.ValueOf(config.HandlerFunction)
		functionName := runtime.FuncForPC(function.Pointer()).Name()
		fmt.Println("Function name:", functionName)
		// get the headline
		headline, err := scraper.ParseFeedItem(config.Url, config.Selector, config.HandlerFunction)
		fmt.Println("Headline:", *headline)
		if err != nil {
			t.Error("Error getting headline:", err)
			return
		}
	}
}