package generators

import (
	"squawkmarketbackend/scraper"
	scraperTypes "squawkmarketbackend/scraper/types"
	"squawkmarketbackend/utils"
)

func GeneratePostmarketMessage() (*string, error) {

	currentESTTime := utils.GetCurrentESTTime()
	sAndPCloseSentence, err := scraper.ScrapeForConfigItem(scraperTypes.YahooSAndPCloseConfig)
	if err != nil {
		return nil, err
	}
	dowCloseSentence, err := scraper.ScrapeForConfigItem(scraperTypes.YahooDowCloseConfig)
	if err != nil {
		return nil, err
	}

	premarketMessage := "Good evening, it's " +
		currentESTTime +
		". This is the Squawk Market postmarket summary. " +
		*sAndPCloseSentence + " " + *dowCloseSentence + " That was the post market summary from Squawk Market, the best real-time & market-wide audio feed."

	return &premarketMessage, nil
}
