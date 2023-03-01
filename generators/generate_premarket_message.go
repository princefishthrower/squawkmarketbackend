package generators

import (
	"squawkmarketbackend/scraper"
	scraperTypes "squawkmarketbackend/scraper/types"
	"squawkmarketbackend/utils"
)

func GeneratePremarketMessage() (*string, error) {

	currentESTTime := utils.GetCurrentESTTime()
	sAndPFuturesSentenceSquawk, err := scraper.ScrapeForConfigItem(scraperTypes.YahooSAndPFuturesConfig)
	if err != nil {
		return nil, err
	}
	if len(sAndPFuturesSentenceSquawk) == 0 {
		return nil, err
	}
	dowFuturesSentenceSquawk, err := scraper.ScrapeForConfigItem(scraperTypes.YahooDowFuturesConfig)
	if err != nil {
		return nil, err
	}
	if len(dowFuturesSentenceSquawk) == 0 {
		return nil, err
	}

	premarketMessage := "Good morning, it's " +
		currentESTTime +
		". This is the Squawk Market premarket summary. " +
		sAndPFuturesSentenceSquawk[0] + " " + dowFuturesSentenceSquawk[0] + " That's it for now, see you at the open."

	return &premarketMessage, nil
}
