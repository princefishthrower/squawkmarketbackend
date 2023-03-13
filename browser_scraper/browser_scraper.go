package browser_scraper

import (
	"log"
	"os"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

func GetNonStaticContent(url string, selector string) (selenium.WebElement, error) {
	// Set up Selenium web driver options
	caps := selenium.Capabilities{}
	chromeCaps := chrome.Capabilities{
		Path: os.Getenv("GOOGLE_CHROME_PATH"),
		Args: []string{
			"--headless",
			"--disable-gpu",
			"--window-size=1920,1080",
		},
	}
	caps.AddChrome(chromeCaps)

	// Create a new WebDriver instance
	wd, err := selenium.NewRemote(caps, "http://localhost:4444/wd/hub")
	if err != nil {
		log.Fatalf("Failed to start Selenium: %v", err)
	}
	defer wd.Quit()

	// Navigate to the url
	err = wd.Get(url)
	if err != nil {
		log.Fatalf("Failed to navigate: %v", err)
	}

	// Wait for the page to fully load (javascript, etc.)
	time.Sleep(5 * time.Second)

	// Get the content by selector
	element, err := wd.FindElement(selenium.ByTagName, selector)
	if err != nil {
		log.Fatalf("Failed to find element: %v", err)
	}

	// return the element
	return element, nil
}
