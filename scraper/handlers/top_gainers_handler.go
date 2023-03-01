package scraper

import (
	"squawkmarketbackend/utils"

	"github.com/gocolly/colly/v2"
)

func TopGainersHandler(squawks []string, url string) func(e *colly.HTMLElement) {
	return func(e *colly.HTMLElement) {
		// in this handler, we get ALL matches and add them to the squawks array
		// loop at all td's in the element
		tdTexts := []string{}
		e.ForEach("td", func(_ int, e *colly.HTMLElement) {
			tdTexts = append(tdTexts, e.Text)
		})

		if len(tdTexts) < 10 {
			return
		}
		if tdTexts[9] == "0.00%" {
			return
		}

		squawk := "New top gainer: " + tdTexts[2] + ", symbol " + tdTexts[1] + ", up " + tdTexts[9] + "."

		// run squawk cleaner utility
		squawk = utils.CleanSquawk(squawk)

		// append squawk to squawks slice
		squawks = append(squawks, squawk)
	}
}
