package browser_scraper

type BrowserScraperConfig struct {
	Url      string
	Selector string
}

var NasdaqBrowserScraperConfig = BrowserScraperConfig{
	Url:      "https://www.nasdaqtrader.com/trader.aspx?id=TradeHalts",
	Selector: "#divTradeHaltResults table tr:nth-child(2)",
}
