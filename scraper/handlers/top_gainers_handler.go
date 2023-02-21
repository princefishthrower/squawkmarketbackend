package scraper

import (
	"squawkmarketbackend/utils"

	"github.com/gocolly/colly/v2"
)

func TopGainersHandler(headline *string, url string) func(e *colly.HTMLElement) {
	return func(e *colly.HTMLElement) {
		// only get the first headline
		if *headline == "" {
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

			*headline = "New top gainer: " + tdTexts[2] + ", symbol " + tdTexts[1] + ", up " + tdTexts[9] + "."

			// run headline cleaner utility
			*headline = utils.CleanHeadline(*headline)
			return
		}
	}
}
