package scraper

import (
	"squawkmarketbackend/utils"

	"github.com/gocolly/colly/v2"
)

func HeadlineHandler(headline *string, url string) func(e *colly.HTMLElement) {
	return func(e *colly.HTMLElement) {
		// only get the first headline
		if *headline == "" {
			*headline = e.Text

			// run headline cleaner utility
			*headline = utils.CleanHeadline(*headline)
			return
		}
	}
}
