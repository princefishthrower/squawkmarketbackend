package tests

import (
	"fmt"
	"reflect"
	"runtime"
	"squawkmarketbackend/scraper"
	scraperTypes "squawkmarketbackend/scraper/types"

	"testing"
)

func TestScrapeFinvizSector(t *testing.T) {

	// here the config is the finviz sector config
	config := scraperTypes.FinvizSectorConfig

	// use reflection to get the function name
	function := reflect.ValueOf(config.HandlerFunction)
	functionName := runtime.FuncForPC(function.Pointer()).Name()
	fmt.Println("Function name:", functionName)
	// get the squawks
	squawks, err := scraper.ScrapeForConfigItem(config)
	if err != nil {
		t.Error("Error getting squawk:", err)
		return
	}
	if len(squawks) == 0 {
		t.Error("Squawks is empty")
		return
	}
	if squawks[0] == "" {
		t.Error("squawks[0] is empty")
		return
	}

	fmt.Println("FIRST SQUAWK: ", squawks[0])
}
