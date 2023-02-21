package utils

import (
	"regexp"
	"strings"
)

func CleanHeadline(headline string) string {
	headline = strings.ReplaceAll(headline, "\n", "")
	headline = strings.ReplaceAll(headline, "\t", "")
	headline = strings.ReplaceAll(headline, "\r", "")
	headline = strings.ReplaceAll(headline, "Exclusive:", "")
	headline = strings.ReplaceAll(headline, "UPDATE 1-", "")
	headline = strings.ReplaceAll(headline, "- sources", "")
	headline = strings.ReplaceAll(headline, "-sources", "")
	headline = strings.ReplaceAll(headline, "US STOCKS-", "US Stocks: ")
	headline = strings.ReplaceAll(headline, "GLOBAL MARKETS-", "Global Markets: ")
	headline = strings.ReplaceAll(headline, "---", " - ")
	m1 := regexp.MustCompile(`[0-9][0-9]:[0-9][0-9][AP]M`)
	replaced := m1.ReplaceAllString(headline, "")
	// and trim any leading and trailing whitespace
	return strings.TrimSpace(replaced)
}
