package parsers

import (
	"context"
	"time"

	"github.com/chromedp/chromedp"
)

// CDiscountParser holds the informations needed to scrap
type CDiscountParser struct {
	Name string
	URL string
	shortURL string
}

type chromeHostJSONVersionResponseBody struct {
	WebSocketDebuggerURL string `json:"webSocketDebuggerUrl"`
}

// NewCDiscountParser returns a new instance of fnac parser
func NewCDiscountParser() CDiscountParser {
	return CDiscountParser{
		Name: "CDiscount",
		URL: "https://www.cdiscount.com/jeux-pc-video-console/ps5/console-ps5-blanche-white-standard-edition-plays/f-1035001-ps5standardwhi.html",
		shortURL: "https://cutt.ly/ljDcHkY",
	}
}

// GetName returns the name of the scraper
func (p CDiscountParser) GetName() string {
	return p.Name
}

// GetURL returns the url of the scraper
func (p CDiscountParser) GetURL() string {
	return p.URL
}

// GetShortURL returns the short URL used for display
func (p CDiscountParser) GetShortURL() string {
	return p.shortURL
}

// IsAvailable check for PS5 availability
func (p CDiscountParser) IsAvailable() (bool, error) {
	chromeURL, err := getRemoteDebuggerURL()

	if err != nil {
		return false, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)

	defer cancel()

	// create allocator context for use with creating a browser context later
	allocatorContext, cancel := chromedp.NewRemoteAllocator(ctx, chromeURL)
	defer cancel()

	// create context
	ctxt, cancel := chromedp.NewContext(allocatorContext)
	defer cancel()

	// run task list
	var res bool
	if err := chromedp.Run(ctxt,
		chromedp.Navigate(p.URL),
		chromedp.WaitReady(".ftBToTop"),
		chromedp.Evaluate(`document.querySelector('.fpTMain') !== null`, &res),
	); err != nil {
		ctxt.Done()
		allocatorContext.Done()

		return false, err
	}

	ctxt.Done()
	allocatorContext.Done()

	if res {
		return true, nil
	}

	return false, nil
}