package scraper

import (
	"squawkmarketbackend/utils"
	"strings"

	"github.com/gocolly/colly/v2"
)

func UnusualVolumeHandler(squawks []string, url string) func(e *colly.HTMLElement) {
	return func(e *colly.HTMLElement) {
		// loop at all td's in the element
		tdTexts := []string{}
		e.ForEach("td", func(_ int, e *colly.HTMLElement) {
			tdTexts = append(tdTexts, e.Text)
		})

		if len(tdTexts) < 11 {
			return
		}
		if tdTexts[9] == "0.00%" {
			return
		}

		var direction = "up"
		if strings.Contains("-", tdTexts[9]) {
			direction = "down"
		}

		humanReadableVolume := utils.LargeNumberToReadingString(tdTexts[10])

		squawk := "Unusual volume: " + tdTexts[2] + ", symbol " + tdTexts[1] + ", has traded over " + humanReadableVolume + " shares, " + direction + " " + tdTexts[9] + " at $" + tdTexts[8] + "."

		// run squawk cleaner utility
		squawk = utils.CleanSquawk(squawk)

		// append squawk to squawks slice
		squawks = append(squawks, squawk)
	}
}
