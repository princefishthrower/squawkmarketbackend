package scraper

import (
	"squawkmarketbackend/utils"

	"github.com/gocolly/colly/v2"
)

func WallStreetJournalHeadlineHandler(squawks []string, url string) func(e *colly.HTMLElement) {
	return func(e *colly.HTMLElement) {

		// only continue if the squawk array is empty
		if len(squawks) != 0 {
			return
		}

		// first we need to get story type - type is class that starts with "WSJTheme--articleType"
		e.ForEach("div[class*='WSJTheme--articleType'] span", func(_ int, articleTypeElement *colly.HTMLElement) {
			// if the class starts with "WSJTheme--articleType" then we have the story type
			articleType := articleTypeElement.Text

			if articleType == "Earnings" || articleType == "Finance" || articleType == "Business" || articleType == "Stocks" || articleType == "U.S. Markets" {
				// use original element to get the headline
				e.ForEach("h2", func(_ int, headlineElement *colly.HTMLElement) {
					headline := headlineElement.Text

					// if the headline is not empty
					if headline != "" {
						squawk := headline + " - " + url
						squawk = utils.CleanSquawk(squawk)
						squawks = append(squawks, squawk)
						return
					}
				})
			}
		})
	}
}
