package scraper

import (
	"squawkmarketbackend/utils"

	"github.com/gocolly/colly/v2"
)

func HeadlineHandler(squawk *string, url string) func(e *colly.HTMLElement) {
	return func(e *colly.HTMLElement) {
		// only get the first headline
		if *squawk == "" {
			*squawk = e.Text

			// run squawk cleaner utility
			*squawk = utils.CleanSquawk(*squawk)
			return
		}
	}
}
