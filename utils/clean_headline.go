package utils

import (
	"regexp"
	"strings"
)

// a last cleaner before the mp3 is generated
func CleanSquawk(squawk string) string {
	squawk = strings.ReplaceAll(squawk, "\n", "")
	squawk = strings.ReplaceAll(squawk, "\t", "")
	squawk = strings.ReplaceAll(squawk, "\r", "")
	squawk = strings.ReplaceAll(squawk, "Exclusive:", "")
	squawk = strings.ReplaceAll(squawk, "Exclusive-", "")
	squawk = strings.ReplaceAll(squawk, "UPDATE 1-", "")
	squawk = strings.ReplaceAll(squawk, "- sources", "")
	squawk = strings.ReplaceAll(squawk, "-sources", "")
	squawk = strings.ReplaceAll(squawk, "US STOCKS-", "US Stocks: ")
	squawk = strings.ReplaceAll(squawk, "GLOBAL MARKETS-", "Global Markets: ")
	squawk = strings.ReplaceAll(squawk, "SNAPSHOT", "")
	squawk = strings.ReplaceAll(squawk, "EXPLAINER-", "")
	squawk = strings.ReplaceAll(squawk, ": Markets Wrap", "")
	squawk = strings.ReplaceAll(squawk, " mln ", " million ")
	squawk = strings.ReplaceAll(squawk, " bln ", " billion ")
	squawk = strings.ReplaceAll(squawk, "Global Markets:", "")
	squawk = strings.ReplaceAll(squawk, "---", " - ")
	m1 := regexp.MustCompile(`[0-9][0-9]:[0-9][0-9][AP]M`)
	replaced := m1.ReplaceAllString(squawk, "")

	// and trim any leading and trailing whitespace
	return strings.TrimSpace(replaced)
}
