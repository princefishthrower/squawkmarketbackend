package fear_and_greed

import (
	"fmt"
	"os"
	fearAndGreedTypes "squawkmarketbackend/fear_and_greed/types"
	"squawkmarketbackend/http_helper"
)

func GetFearAndGreedSquawk() (string, error) {
	// get the fear and greed index
	fearAndGreedIndex, err := GetFearAndGreedIndex()
	if err != nil {
		return "", err
	}

	// get the squawk
	squawk := fmt.Sprintf("The CNN Fear and Greed Index is currently at %d, representing '%s'.", fearAndGreedIndex.Fgi.Now.Value, fearAndGreedIndex.Fgi.Now.ValueText)

	return squawk, nil
}

func GetFearAndGreedIndex() (*fearAndGreedTypes.FearAndGreedResponse, error) {

	requestUrl := "https://fear-and-greed-index.p.rapidapi.com/v1/fgi"

	// the headers to pass
	headers := map[string]string{
		"X-RapidAPI-Key":  os.Getenv("FEAR_AND_GREED_KEY"),
		"X-RapidAPI-Host": os.Getenv("FEAR_AND_GREED_HOST"),
	}

	response := fearAndGreedTypes.FearAndGreedResponse{}

	// call the function
	response, err := http_helper.MakeHTTPRequestGeneric(requestUrl, "GET", headers, nil, nil, response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
