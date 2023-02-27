package generators

import (
	"fmt"
	"squawkmarketbackend/pdfdownloader"
	"squawkmarketbackend/pdfextractor"
	"squawkmarketbackend/utils"
	"strings"
)

var countStrings = []string{"First", "Second", "Third", "Fourth"}

func GenerateFedMeetingMinutesSummary() (*string, error) {
	fileName := "fomcminutes20230201.pdf"
	// download using pdfdownloader
	err := pdfdownloader.DownloadPdf("https://www.federalreserve.gov/monetarypolicy/files/fomcminutes20230201.pdf", fileName)
	if err != nil {
		return nil, err
	}

	// convert to text using pdfextractor
	text, err := pdfextractor.GetPdfText(fileName)
	if err != nil {
		return nil, err
	}

	// parse text - break into sentences
	sentences := strings.Split(*text, ".")

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

	return &message, nil
}
