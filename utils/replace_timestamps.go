package utils

import (
	"regexp"
	"strings"
)

func ReplaceTimestamps(text string) string {
	m1 := regexp.MustCompile(`[0-9][0-9]:[0-9][0-9][AP]M`)
	replaced := m1.ReplaceAllString(text, "")
	// and trim any leading and trailing whitespace
	return strings.TrimSpace(replaced)
}
