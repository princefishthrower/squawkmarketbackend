package scraper

import (
	"strings"

	"github.com/gocolly/colly/v2"
)

func FOMCMeetingMinutesHandler(squawk *string, url string) func(e *colly.HTMLElement) {
	return func(e *colly.HTMLElement) {
		// just return the whole text content - it is parsed differently further downstream
		if *squawk == "" {
			*squawk = e.Text
			// remove newlines, replace with spaces
			*squawk = strings.ReplaceAll(*squawk, "\n", " ")
			return
		}
	}
}
