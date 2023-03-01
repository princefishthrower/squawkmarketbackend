package generators

import (
	"squawkmarketbackend/scraper"
	scraperTypes "squawkmarketbackend/scraper/types"
	"squawkmarketbackend/utils"
)

func GeneratePostmarketMessage() (*string, error) {

	currentESTTime := utils.GetCurrentESTTime()
	sAndPCloseSentenceSquawks, err := scraper.ScrapeForConfigItem(scraperTypes.YahooSAndPCloseConfig)
	if err != nil {
		return nil, err
	}
	if len(sAndPCloseSentenceSquawks) == 0 {
		return nil, err
	}
	dowCloseSentenceSquawks, err := scraper.ScrapeForConfigItem(scraperTypes.YahooDowCloseConfig)
	if err != nil {
		return nil, err
	}
	if len(dowCloseSentenceSquawks) == 0 {
		return nil, err
	}

	premarketMessage := "Good evening, it's " +
		currentESTTime +
		". This is the Squawk Market postmarket summary. " +
		sAndPCloseSentenceSquawks[0] + " " + dowCloseSentenceSquawks[0] + " That was the post market summary from Squawk Market, the best real-time & market-wide audio feed."

	return &premarketMessage, nil
}
