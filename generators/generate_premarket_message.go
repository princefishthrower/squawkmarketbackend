package generators

import (
	"squawkmarketbackend/scraper"
	scraperTypes "squawkmarketbackend/scraper/types"
	"squawkmarketbackend/utils"
)

func GeneratePremarketMessage() (*string, error) {

	currentESTTime := utils.GetCurrentESTTime()
	sAndPFuturesSentence, err := scraper.ScrapeForConfigItem(scraperTypes.YahooSAndPFuturesConfig)
	if err != nil {
		return nil, err
	}
	dowFuturesSentence, err := scraper.ScrapeForConfigItem(scraperTypes.YahooDowFuturesConfig)
	if err != nil {
		return nil, err
	}

	premarketMessage := "Good morning, it's " +
		currentESTTime +
		". This is the Squawk Market premarket summary. " +
		*sAndPFuturesSentence + " " + *dowFuturesSentence + " That's it for now, see you at the open."

	return &premarketMessage, nil
}
