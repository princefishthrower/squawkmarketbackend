package headlines

type Headline struct {
	ID        int    `json:"id"`
	CreatedAt string `json:"created_at"`
	Headline  string `json:"headline"`
	Mp3Data   []byte `json:"mp3data"`
}

type HeadlineConfig struct {
	Url      string
	Selector string
}

// define a const config of type slice HeadlineConfig for each of the news sources
var MarketWatchConfig = HeadlineConfig{
	Url:      "https://marketwatch.com/latest-news",
	Selector: "h3.article__headline",
}

var WallStreetJournalConfig = HeadlineConfig{
	Url:      "https://wsj.com/news/latest-headlines",
	Selector: "h2[class*='WSJTheme--headline']",
}

var ReutersConfig = HeadlineConfig{
	Url:      "https://reuters.com/markets/us/",
	Selector: "a[class*='heading__heading'] span",
}

var YahooConfig = HeadlineConfig{
	Url:      "https://finance.yahoo.com/news",
	Selector: "a.js-content-viewer",
}

var FinvizConfig = HeadlineConfig{
	Url:      "https://finviz.com/news.ashx",
	Selector: "tr.nn",
}

// now define the slice of configs
var HeadlineConfigs = []HeadlineConfig{
	MarketWatchConfig,
	WallStreetJournalConfig,
	ReutersConfig,
	YahooConfig,
	FinvizConfig,
}
