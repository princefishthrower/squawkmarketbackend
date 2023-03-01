package scraper

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
)

func EconomicCalendarHandler(squawk *string, url string) func(e *colly.HTMLElement) {
	return func(e *colly.HTMLElement) {
		// log out element text
		fmt.Println(e.DOM.Text())

		// get all the td elements
		tds := e.DOM.Find("td")

		// if the second td element includes "USDs" AND the fifth td is not empty
		if *squawk == "" && strings.Contains(tds.Eq(1).Text(), "USD") && notEmpty(tds.Eq(4).Text()) {
			// build the squawk phrase
			*squawk = "The economic print for " + tds.Eq(3).Text() + " is just in at " + tds.Eq(4).Text() + ". Expected value was  " + tds.Eq(5).Text() + ", previous value was " + tds.Eq(6).Text() + "."

			// clean up the squawk - remove newlines
			*squawk = strings.ReplaceAll(*squawk, "\n", " ")
		}
	}
}

func notEmpty(actual string) bool {
	if actual == "" {
		return false
	}
	if strings.Contains(actual, "\u00a0") {
		return false
	}
	return true
}
