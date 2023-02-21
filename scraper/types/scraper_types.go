package scraper

import (
	handlers "squawkmarketbackend/scraper/handlers"

	"github.com/gocolly/colly/v2"
)

type Headline struct {
	ID        int    `json:"id"`
	CreatedAt string `json:"created_at"`
	Headline  string `json:"headline"`
	Mp3Data   []byte `json:"mp3data"`
}

type HeadlineConfig struct {
	Url             string
	Selector        string
	HandlerFunction func(*string, string) func(*colly.HTMLElement)
}

// define a const config of type slice HeadlineConfig for each of the news sources
var MarketWatchConfig = HeadlineConfig{
	Url:             "https://marketwatch.com/latest-news",
	Selector:        "h3.article__headline",
	HandlerFunction: handlers.HeadlineHandler,
}

var WallStreetJournalConfig = HeadlineConfig{
	Url:             "https://wsj.com/news/latest-headlines",
	Selector:        "h2[class*='WSJTheme--headline']",
	HandlerFunction: handlers.HeadlineHandler,
}

var ReutersConfig = HeadlineConfig{
	Url:             "https://reuters.com/markets/us/",
	Selector:        "a[class*='heading__heading'] span",
	HandlerFunction: handlers.HeadlineHandler,
}

var YahooConfig = HeadlineConfig{
	Url:             "https://finance.yahoo.com/news",
	Selector:        "a.js-content-viewer",
	HandlerFunction: handlers.HeadlineHandler,
}

var FinvizConfig = HeadlineConfig{
	Url:             "https://finviz.com/news.ashx",
	Selector:        "tr.nn",
	HandlerFunction: handlers.HeadlineHandler,
}

var FinvizTopGainersConfig = HeadlineConfig{
	Url:             "https://finviz.com/screener.ashx?s=ta_topgainers",
	Selector:        "table .table-light tr:nth-child(2)",
	HandlerFunction: handlers.TopGainersHandler,
}

var FinvizTopLosersConfig = HeadlineConfig{
	Url:             "https://finviz.com/screener.ashx?s=ta_toplosers",
	Selector:        "table .table-light tr:nth-child(2)",
	HandlerFunction: handlers.TopLosersHandler,
}

var FinvizNewHighConfig = HeadlineConfig{
	Url:             "https://finviz.com/screener.ashx?s=ta_newhigh",
	Selector:        "table .table-light tr:nth-child(2)",
	HandlerFunction: handlers.NewHighHandler,
}

var FinvizNewLowConfig = HeadlineConfig{
	Url:             "https://finviz.com/screener.ashx?s=ta_newlow",
	Selector:        "table .table-light tr:nth-child(2)",
	HandlerFunction: handlers.NewLowHandler,
}

var FinvizOverboughtConfig = HeadlineConfig{
	Url:             "https://finviz.com/screener.ashx?s=ta_overbought&o=-change",
	Selector:        "table .table-light tr:nth-child(2)",
	HandlerFunction: handlers.OverboughtHandler,
}

var FinvizOversoldConfig = HeadlineConfig{
	Url:             "https://finviz.com/screener.ashx?s=ta_oversold&o=change",
	Selector:        "table .table-light tr:nth-child(2)",
	HandlerFunction: handlers.OversoldHandler,
}

var FinvizUnusualVolumeConfig = HeadlineConfig{
	Url:             "https://finviz.com/screener.ashx?s=ta_unusualvolume&o=-volume",
	Selector:        "table .table-light tr:nth-child(2)",
	HandlerFunction: handlers.UnusualVolumeHandler,
}

var FinvizMostVolatileConfig = HeadlineConfig{
	Url:             "https://finviz.com/screener.ashx?s=ta_mostvolatile",
	Selector:        "table .table-light tr:nth-child(2)",
	HandlerFunction: handlers.MostVolatileHandler,
}

// now define the slice of configs
var HeadlineConfigs = []HeadlineConfig{
	MarketWatchConfig,
	WallStreetJournalConfig,
	ReutersConfig,
	YahooConfig,
	FinvizConfig,
	FinvizTopGainersConfig,
	FinvizTopLosersConfig,
	FinvizNewHighConfig,
	FinvizNewLowConfig,
	FinvizOverboughtConfig,
	FinvizOversoldConfig,
	FinvizUnusualVolumeConfig,
	FinvizMostVolatileConfig,
}
