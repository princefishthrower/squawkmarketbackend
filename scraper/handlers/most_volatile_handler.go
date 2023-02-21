package scraper

import (
	"squawkmarketbackend/utils"
	"strings"

	"github.com/gocolly/colly/v2"
)

func MostVolatileHandler(headline *string, url string) func(e *colly.HTMLElement) {
	return func(e *colly.HTMLElement) {
		// only get the first headline
		if *headline == "" {
			// loop at all td's in the element
			tdTexts := []string{}
			e.ForEach("td", func(_ int, e *colly.HTMLElement) {
				tdTexts = append(tdTexts, e.Text)
			})

			if len(tdTexts) < 11 {
				return
			}

			var direction = "up"
			if strings.Contains("-", tdTexts[9]) {
				direction = "down"
			}
			if tdTexts[9] == "0.00%" {
				return
			}

			*headline = "Most volatile: " + tdTexts[2] + ", symbol " + tdTexts[1] + ", has had a large high / low trading range " + direction + " " + tdTexts[9] + " at $" + tdTexts[8] + "."

			// run headline cleaner utility
			*headline = utils.CleanHeadline(*headline)

			return
		}
	}
}
