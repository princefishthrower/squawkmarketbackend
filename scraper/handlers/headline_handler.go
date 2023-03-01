package scraper

import (
	"squawkmarketbackend/utils"

	"github.com/gocolly/colly/v2"
)

func HeadlineHandler(squawks []string, url string) func(e *colly.HTMLElement) {
	return func(e *colly.HTMLElement) {
		// only get the first headline
		if len(squawks) == 0 {
			squawk := e.Text

			// run squawk cleaner utility
			squawk = utils.CleanSquawk(squawk)

			// append squawk to squawks slice
			squawks = append(squawks, squawk)
			return
		}
	}
}
