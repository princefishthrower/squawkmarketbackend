package browser_scraper

import (
	"squawkmarketbackend/browser_scraper"
	browserScraperTypes "squawkmarketbackend/browser_scraper/types"
	"strings"

	"github.com/tebeka/selenium"
)

type NasdaqTradingHaltCode struct {
	Code        string
	Description string
}

var config = []NasdaqTradingHaltCode{
	{"T1", "Halt - News Pending"},
	{"T2", "Halt - News Released"},
	{"T5", "Single Stock Trading Pause in Effect"},
	{"T6", "Halt - Extraordinary Market Activity"},
	{"T8", "Halt - Exchange-Traded-Fund (ETF)"},
	{"T12", "Halt - Additional Information Requested by NASDAQ"},
	{"H4", "Halt - Non-compliance"},
	{"H9", "Halt - Not Current"},
	{"H10", "Halt - SEC Trading Suspension"},
	{"H11", "Halt - Regulatory Concern"},
	{"O1", "Operations Halt, Contact Market Operations"},
	{"IPO1", "IPO Issue not yet Trading"},
	{"IPOQ", "IPO security released for quotation"},
	{"IPOE", "IPO security - positioning window extension"},
	{"M1", "Corporate Action"},
	{"M2", "Quotation Not Available"},
	{"MWC0", "Market Wide Circuit Breaker Halt - Carry over from previous day"},
	{"MWC1", "Market Wide Circuit Breaker Halt - Level 1"},
	{"MWC2", "Market Wide Circuit Breaker Halt - Level 2"},
	{"MWC3", "Market Wide Circuit Breaker Halt - Level 3"},
	{"MWCQ", "Market Wide Circuit Breaker Resumption"},
	{"M", "Volatility Trading Pause"},
	{"D", "Security deletion from NASDAQ / CQS"},
	{"R1", "New Issue Available"},
	{"R2", "Issue Available"},
	{"R4", "Qualifications Issues Reviewed/Resolved; Quotations/Trading to Resume"},
	{"R9", "Filing Requirements Satisfied/Resolved; Quotations/Trading To Resume"},
	{"C3", "Issuer News Not Forthcoming; Quotations/Trading To Resume"},
	{"C4", "Qualifications Halt ended; maint. req. met; Resume"},
	{"C9", "Qualifications Halt Concluded; Filings Met; Quotes/Trades To Resume"},
	{"C11", "Trade Halt Concluded By Other Regulatory Auth,; Quotes/Trades Resume"},
	{"H4", "Halt - Non-compliance"},
	{"H9", "Halt - Not Current"},
	{"H10", "Halt - SEC Trading Suspension"},
	{"H11", "Halt - Regulatory Concern"},
	{"LUDP", "Volatility Trading Pause"},
	{"LUDS", "Volatility Trading Pause - Straddle Condition"},
}

func GetLatestNasdaqTradingHalt() (*string, error) {
	// here the config is the finviz sector config
	element, err := browser_scraper.GetNonStaticContent(browserScraperTypes.NasdaqBrowserScraperConfig.Url, browserScraperTypes.NasdaqBrowserScraperConfig.Selector)
	if err != nil {
		return nil, err
	}

	// Find all td elements within the table element.
	tdElements, err := element.FindElements(selenium.ByTagName, "td")
	if err != nil {
		return nil, err
	}

	// Get texts of each td
	tdTexts := []string{}
	for _, td := range tdElements {
		// Do something with the td element, e.g. get its text content.
		tdText, err := td.Text()
		if err != nil {
			return nil, err
		}
		tdTexts = append(tdTexts, tdText)
	}

	// build squawk text from tds
	symbol := tdTexts[3]
	market := tdTexts[4]
	reasonString := tdTexts[5]

	// trim whitespace on reasonString
	reasonString = strings.TrimSpace(reasonString)

	// split tdText[5] into codes
	reasons := []string{}
	for _, code := range strings.Split(reasonString, " ") {
		// use type to look up description
		for _, config := range config {
			if config.Code == code {
				reasons = append(reasons, config.Description)
			}
		}
	}

	// combine reasons into one string
	combinedReasonString := strings.Join(reasons, ", ")

	squawk := market + "trading halt on " + symbol + " for " + combinedReasonString + "."

	return &squawk, nil

}
