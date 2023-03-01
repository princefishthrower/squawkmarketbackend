package scraper

import (
	"strings"

	"github.com/gocolly/colly/v2"
)

func YahooFuturesHandler(squawks []string, url string) func(e *colly.HTMLElement) {
	return func(e *colly.HTMLElement) {
		// only get the first element found
		if len(squawks) == 0 {
			// get the ariaLabel of the element
			squawk := e.Attr("aria-label")

			// replace the word "has" with "have"
			squawk = strings.ReplaceAll(squawk, "has", "have")

			// add period to the end of the the sentence
			squawk = squawk + "."

			// append squawk to squawks slice
			squawks = append(squawks, squawk)
			return
		}
	}
}
