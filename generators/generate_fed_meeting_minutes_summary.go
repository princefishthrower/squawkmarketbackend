package generators

import (
	"fmt"
	"squawkmarketbackend/openai"
	"squawkmarketbackend/scraper"
	scraperTypes "squawkmarketbackend/scraper/types"
	"squawkmarketbackend/utils"
	"strings"
)

var countStrings = []string{"First", "Second", "Third", "Fourth"}

func GenerateFedMeetingMinutesSummary() (*string, error) {
	texts, err := scraper.ScrapeForConfigItem(scraperTypes.FOMCMeetingMinutesConfig)
	if err != nil {
		return nil, err
	}
	if len(texts) == 0 {
		return nil, err
	}

	// parse text - break into sentences
	sentences := strings.Split(texts[0], ".")

	// for each sentence, check if it contains "basis points"
	matchedSentences := utils.SearchSentences(sentences, "basis points")

	message := "Initial summary of the Fed meeting minutes concerning basis points: ..."
	if len(matchedSentences) == 1 {
		// combine the matches into a single string
		message += strings.Join(matchedSentences, " ") + "..."
	}
	if len(matchedSentences) > 1 {
		// loop at matched sentences, prefixing with "____ excerpt:"
		for i, sentence := range matchedSentences {
			message += fmt.Sprintf("%s excerpt: %s...", countStrings[i], sentence)
		}
	}

	// now open AI to generate a summary
	summary, err := openai.AskGPT("Can you summarize the following to just 4-5 sentences? \n\n" + message)
	if err != nil {
		return nil, err
	}

	return summary, nil
}
