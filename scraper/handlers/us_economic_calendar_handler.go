package scraper

import (
	"squawkmarketbackend/utils"
	"strings"

	"github.com/gocolly/colly/v2"
)

func USEconomicCalendarHandler(squawks *[]string, url string) func(e *colly.HTMLElement) {
	return func(e *colly.HTMLElement) {
		// log out element text
		// fmt.Println(e.DOM.Text())

		// get all the td elements
		tds := e.DOM.Find("td")

		// if the second td element includes "USDs" AND the fifth td is not empty
		if len(*squawks) == 0 && strings.Contains(tds.Eq(1).Text(), "USD") && utils.NotEmpty(tds.Eq(4).Text()) {
			// build the squawk phrase
			squawk := "U.S. " + tds.Eq(3).Text() + " is " + tds.Eq(4).Text()

			if utils.NotEmpty(tds.Eq(5).Text()) {
				squawk += ", expected " + tds.Eq(5).Text()
			}

			if utils.NotEmpty(tds.Eq(6).Text()) {
				squawk += ", previously " + tds.Eq(6).Text()
			}

			squawk += "."

			// clean up the squawk - remove newlines
			squawk = strings.ReplaceAll(squawk, "\n", " ")

			// append squawk to squawks slice
			*squawks = append(*squawks, squawk)
		}
	}
}
