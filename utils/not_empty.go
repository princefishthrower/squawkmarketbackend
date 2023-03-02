package utils

import "strings"

func NotEmpty(actual string) bool {
	if actual == "" {
		return false
	}
	if strings.Contains(actual, "\u00a0") {
		return false
	}
	return true
}
