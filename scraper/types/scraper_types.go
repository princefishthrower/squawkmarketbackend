package scraper

import (
	handlers "squawkmarketbackend/scraper/handlers"

	"github.com/gocolly/colly/v2"
)

// IMPORTANT! ADD SITES TO THIS SLICE
var AllowedDomains = []string{
	"finviz.com",
	"www.finviz.com",
	"finance.yahoo.com",
	"www.finance.yahoo.com",
	"marketwatch.com",
	"www.marketwatch.com",
	"reuters.com",
	"www.reuters.com",
	"wsj.com",
	"www.wsj.com",
	"cryptonews.com",
	"coindesk.com",
	"www.coindesk.com",
	"investing.com",
	"www.investing.com",
	"forbes.com",
	"www.forbes.com",
	"economist.com",
	"www.economist.com",
	"bloomberg.com",
	"www.bloomberg.com",
	"www.federalreserve.gov",
	"benzinga.com",
	"www.benzinga.com",
	"fortune.com",
	"www.fortune.com",
}

type ScrapingConfig struct {
	Url             string
	Selector        string
	HandlerFunction func(*[]string, string) func(*colly.HTMLElement)
	FeedName        string
	InsertThreshold float64
}

// define a const config of type slice ScrapingConfig for each of the news sources
var MarketWatchNewsConfig = ScrapingConfig{
	Url:             "https://marketwatch.com/latest-news",
	Selector:        "h3.article__headline",
	HandlerFunction: handlers.HeadlineHandler,
	FeedName:        "market-wide",
	InsertThreshold: 0.75,
}

var WallStreetJournalNewsConfig = ScrapingConfig{
	Url:             "https://wsj.com/news/latest-headlines",
	Selector:        "article[class*='WSJTheme--story']",
	HandlerFunction: handlers.WallStreetJournalHeadlineHandler,
	FeedName:        "market-wide",
	InsertThreshold: 0.75,
}

var ReutersNewsConfig = ScrapingConfig{
	Url:             "https://reuters.com/markets/us/",
	Selector:        "a[class*='heading__heading'] span",
	HandlerFunction: handlers.HeadlineHandler,
	FeedName:        "market-wide",
	InsertThreshold: 0.75,
}

var YahooNewsConfig = ScrapingConfig{
	Url:             "https://finance.yahoo.com/news",
	Selector:        "a.js-content-viewer",
	HandlerFunction: handlers.HeadlineHandler,
	FeedName:        "market-wide",
	InsertThreshold: 0.75,
}

var FinvizNewsConfig = ScrapingConfig{
	Url:             "https://finviz.com/news.ashx",
	Selector:        "tr.nn",
	HandlerFunction: handlers.HeadlineHandler,
	FeedName:        "market-wide",
	InsertThreshold: 0.75,
}

var CryptonewsNewsConfig = ScrapingConfig{
	Url:             "https://cryptonews.com/",
	Selector:        "a.article__title",
	HandlerFunction: handlers.HeadlineHandler,
	FeedName:        "crypto",
	InsertThreshold: 0.75,
}

var CoinDeskNewsConfig = ScrapingConfig{
	Url:             "https://coindesk.com/",
	Selector:        "div[class*='live-wirestyles__Title'] a",
	HandlerFunction: handlers.HeadlineHandler,
	FeedName:        "crypto",
	InsertThreshold: 0.75,
}

var ForbesNewsConfig = ScrapingConfig{
	Url:             "https://www.forbes.com/news",
	Selector:        "h3 a.stream-item__title",
	HandlerFunction: handlers.HeadlineHandler,
	FeedName:        "crypto",
	InsertThreshold: 0.75,
}

var EconomistBusinessNewsConfig = ScrapingConfig{
	Url:             "https://www.economist.com/business",
	Selector:        "h3 a",
	HandlerFunction: handlers.HeadlineHandler,
	FeedName:        "market-wide",
	InsertThreshold: 0.75,
}

var EconomistFinanceNewsConfig = ScrapingConfig{
	Url:             "https://www.economist.com/finance-and-economics",
	Selector:        "h3 a",
	HandlerFunction: handlers.HeadlineHandler,
	FeedName:        "market-wide",
	InsertThreshold: 0.75,
}

var BloombergMarketNewsConfig = ScrapingConfig{
	Url:             "https://www.bloomberg.com/markets",
	Selector:        ".single-story-module__headline-link",
	HandlerFunction: handlers.HeadlineHandler,
	FeedName:        "market-wide",
	InsertThreshold: 0.75,
}

var FortuneNewsConfig = ScrapingConfig{
	Url:             "https://fortune.com/",
	Selector:        "ul li div a .titleLink",
	HandlerFunction: handlers.HeadlineHandler,
	FeedName:        "market-wide",
	InsertThreshold: 0.75,
}

var BenzingaNewsConfig = ScrapingConfig{
	Url:             "https://www.benzinga.com/recent",
	Selector:        ".post-title span",
	HandlerFunction: handlers.HeadlineHandler,
	FeedName:        "market-wide",
	InsertThreshold: 0.75,
}

var FinvizTopGainersConfig = ScrapingConfig{
	Url:             "https://finviz.com/screener.ashx?s=ta_topgainers&o=-change",
	Selector:        "table .table-light tr:nth-child(2)",
	HandlerFunction: handlers.TopGainersHandler,
	FeedName:        "top-gainers",
	InsertThreshold: 0.90,
}

var FinvizTopLosersConfig = ScrapingConfig{
	Url:             "https://finviz.com/screener.ashx?s=ta_toplosers&o=change",
	Selector:        "table .table-light tr:nth-child(2)",
	HandlerFunction: handlers.TopLosersHandler,
	FeedName:        "top-losers",
	InsertThreshold: 0.90,
}

var FinvizNewHighConfig = ScrapingConfig{
	Url:             "https://finviz.com/screener.ashx?s=ta_newhigh&o=-change",
	Selector:        "table .table-light tr:nth-child(2)",
	HandlerFunction: handlers.NewHighHandler,
	FeedName:        "new-highs",
	InsertThreshold: 0.90,
}

var FinvizNewLowConfig = ScrapingConfig{
	Url:             "https://finviz.com/screener.ashx?s=ta_newlow&o=change",
	Selector:        "table .table-light tr:nth-child(2)",
	HandlerFunction: handlers.NewLowHandler,
	FeedName:        "new-lows",
	InsertThreshold: 0.90,
}

var FinvizOverboughtConfig = ScrapingConfig{
	Url:             "https://finviz.com/screener.ashx?s=ta_overbought&o=-change",
	Selector:        "table .table-light tr:nth-child(2)",
	HandlerFunction: handlers.OverboughtHandler,
	FeedName:        "overbought",
	InsertThreshold: 0.90,
}

var FinvizOversoldConfig = ScrapingConfig{
	Url:             "https://finviz.com/screener.ashx?s=ta_oversold&o=change",
	Selector:        "table .table-light tr:nth-child(2)",
	HandlerFunction: handlers.OversoldHandler,
	FeedName:        "oversold",
	InsertThreshold: 0.90,
}

var FinvizUnusualVolumeConfig = ScrapingConfig{
	Url:             "https://finviz.com/screener.ashx?s=ta_unusualvolume&o=-volume",
	Selector:        "table .table-light tr:nth-child(2)",
	HandlerFunction: handlers.UnusualVolumeHandler,
	FeedName:        "unusual-trading-volume",
	InsertThreshold: 0.90,
}

var FinvizMostVolatileConfig = ScrapingConfig{
	Url:             "https://finviz.com/screener.ashx?s=ta_mostvolatile",
	Selector:        "table .table-light tr:nth-child(2)",
	HandlerFunction: handlers.MostVolatileHandler,
	FeedName:        "most-volatile",
	InsertThreshold: 0.90,
}

var YahooSAndPFuturesConfig = ScrapingConfig{
	Url:             "https://finance.yahoo.com",
	Selector:        "a[title~='S&P']",
	HandlerFunction: handlers.YahooFuturesHandler,
	FeedName:        "futures",
	InsertThreshold: 0.0,
}

var YahooDowFuturesConfig = ScrapingConfig{
	Url:             "https://finance.yahoo.com",
	Selector:        "a[title~='Dow']",
	HandlerFunction: handlers.YahooFuturesHandler,
	FeedName:        "futures",
	InsertThreshold: 0.0,
}

var YahooSAndPCloseConfig = ScrapingConfig{
	Url:             "https://finance.yahoo.com",
	Selector:        "a[title~='S&P']",
	HandlerFunction: handlers.YahooFuturesHandler,
	FeedName:        "market-wide",
	InsertThreshold: 0.0,
}

var YahooDowCloseConfig = ScrapingConfig{
	Url:             "https://finance.yahoo.com",
	Selector:        "a[title~='Dow']",
	HandlerFunction: handlers.YahooFuturesHandler,
	FeedName:        "market-wide",
	InsertThreshold: 0.0,
}

var EconomicCalendarConfig = ScrapingConfig{
	Url:             "https://investing.com/economic-calendar/",
	Selector:        "table tr",
	HandlerFunction: handlers.EconomicCalendarHandler,
	FeedName:        "economic-prints",
	InsertThreshold: 0.0,
}

var FOMCMeetingMinutesConfig = ScrapingConfig{
	Url:             "https://www.federalreserve.gov/monetarypolicy/fomcminutes20230201.htm",
	Selector:        "#article",
	HandlerFunction: handlers.FOMCMeetingMinutesHandler,
	FeedName:        "economic-prints",
	InsertThreshold: 0.0,
}

// now define the slice of configs
var ScrapingConfigs = []ScrapingConfig{
	MarketWatchNewsConfig,
	WallStreetJournalNewsConfig,
	ReutersNewsConfig,
	YahooNewsConfig,
	FinvizNewsConfig,
	CryptonewsNewsConfig,
	CoinDeskNewsConfig,
	FinvizTopGainersConfig,
	FinvizTopLosersConfig,
	FinvizNewHighConfig,
	FinvizNewLowConfig,
	FinvizOverboughtConfig,
	FinvizOversoldConfig,
	FinvizUnusualVolumeConfig,
	FinvizMostVolatileConfig,
	EconomicCalendarConfig,
}
