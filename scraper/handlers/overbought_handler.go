package scraper

import (
	"squawkmarketbackend/utils"

	"github.com/gocolly/colly/v2"
)

func OverboughtHandler(squawks []string, url string) func(e *colly.HTMLElement) {
	return func(e *colly.HTMLElement) {

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

		squawk := "Overbought according to RSI(14): " + tdTexts[2] + ", symbol " + tdTexts[1] + ", at $" + tdTexts[8] + ", up " + tdTexts[9] + "."

		// run squawk cleaner utility
		squawk = utils.CleanSquawk(squawk)

		// append squawk to squawks slice
		squawks = append(squawks, squawk)
	}
}
