package scraper

import (
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
)

func SectorHandler(squawks *[]string, url string) func(e *colly.HTMLElement) {
	return func(e *colly.HTMLElement) {
		// get all rows from the sector table
		rows := []colly.HTMLElement{}
		e.ForEach("tr", func(_ int, e *colly.HTMLElement) {
			rows = append(rows, *e)
		})

		// loop over all rows
		for _, row := range rows {
			// get all columns from the row
			cols := []string{}
			row.ForEach("td", func(_ int, e *colly.HTMLElement) {
				cols = append(cols, e.Text)
			})
			if len(cols) < 10 {
				continue
			}
			var change = cols[9]
			var direction = "up"

			// split change at decimal point, convert to int, and only squawk if it is greater than 1
			changeSplit := strings.Split(change, ".")
			if len(changeSplit) < 2 {
				continue
			}
			changeInt, err := strconv.Atoi(changeSplit[0])
			if err != nil {
				continue
			}
			// if absolute value of changeInt is less than 1 then continue
			if changeInt > -1 && changeInt < 1 {
				continue
			}

			// if changeInt is negative then change direction to down and remove the negative sign
			if changeInt < 0 {
				direction = "down"
				// remove the negative sign
				change = strings.ReplaceAll(change, "-", "")
			}
			sector := cols[1]

			// convert to lower case
			sector = strings.ToLower(sector)

			squawk := "The " + sector + " sector is " + direction + " " + change + "."
			// append squawk to squawks slice
			*squawks = append(*squawks, squawk)
		}
	}
}
