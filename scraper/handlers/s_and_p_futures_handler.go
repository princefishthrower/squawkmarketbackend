package scraper

import (
	"strings"

	"github.com/gocolly/colly/v2"
)

func YahooFuturesHandler(squawk *string, url string) func(e *colly.HTMLElement) {
	return func(e *colly.HTMLElement) {
		// only get the first element found
		if *squawk == "" {
			// get the ariaLabel of the element
			*squawk = e.Attr("aria-label")

			// replace the word "has" with "have"
			*squawk = strings.ReplaceAll(*squawk, "has", "have")

			// add period to the end of the the sentance
			*squawk = *squawk + "."

			return
		}
	}
}
