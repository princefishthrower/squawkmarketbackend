package utils

import (
	"math"
	"strings"
)

func SearchSentences(sentences []string, keyword string) []string {
	var result []string

	// iterate through each sentence in the array
	for i, sentence := range sentences {
		// if the sentence contains the keyword, add the 3 sentences on each side to the result array
		if strings.Contains(sentence, keyword) {
			startIndex := int(math.Max(0, float64(i-3)))
			endIndex := int(math.Min(float64(len(sentences)), float64(i+4)))
			result = append(result, strings.Join(sentences[startIndex:endIndex], " "))
		}
	}

	return result
}
