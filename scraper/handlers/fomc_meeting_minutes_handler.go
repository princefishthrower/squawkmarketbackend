package scraper

import (
	"strings"

	"github.com/gocolly/colly/v2"
)

func FOMCMeetingMinutesHandler(squawks []string, url string) func(e *colly.HTMLElement) {
	return func(e *colly.HTMLElement) {
		// just return the whole text content - it is parsed differently further downstream
		// should be only a single match anyway since we match ID "#article"
		if len(squawks) == 0 {
			squawk := e.Text
			// remove newlines, replace with spaces
			squawk = strings.ReplaceAll(squawk, "\n", " ")

			// append squawk to squawks slice
			squawks = append(squawks, squawk)
		}
	}
}
