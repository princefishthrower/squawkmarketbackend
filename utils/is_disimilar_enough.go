package utils

import (
	"github.com/adrg/strutil"
	"github.com/adrg/strutil/metrics"
)

// AreStringsTooSimilar compares two strings and returns true if they are too similar (over threshold)
func AreStringsTooSimilar(s1 string, s2 string, threshold float64) bool {
	similarity := strutil.Similarity(s1, s2, metrics.NewLevenshtein())
	return similarity > threshold
}
