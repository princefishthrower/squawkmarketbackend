package textextractor

import "strings"

func ExtractSubstring(input, target string, numWords int) []string {
	// Split the input string into words.
	words := strings.Fields(input)

	// Find the index of the target word.
	targetIndex := -1
	for i, word := range words {
		if word == target {
			targetIndex = i
			break
		}
	}

	// If the target word was not found, return an empty result.
	if targetIndex == -1 {
		return []string{}
	}

	// Calculate the start and end indexes of the substring.
	startIndex := targetIndex - numWords
	if startIndex < 0 {
		startIndex = 0
	}
	endIndex := targetIndex + numWords + 1
	if endIndex > len(words) {
		endIndex = len(words)
	}

	// Extract the substring.
	result := strings.Join(words[startIndex:endIndex], " ")

	return []string{result}
}
