package parsers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"github.com/jeromewir/stockalerts/config"
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
		URL: "https://www.cdiscount.com/jeux-pc-video-console/ps5/console-ps5/l-1035001.html#_his_",
		shortURL: "https://cutt.ly/WjziXL5",
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

func (p CDiscountParser) getRemoteDebuggerURL() (string, error) {
	chromeHost := config.GetChromeHost()
	chromePort := config.GetChromePort()

	ips, err := net.LookupIP(chromeHost)

	if err != nil {
		return "", err
	}

	if len(ips) > 0 {
		chromeHost = ips[0].String()
	}

	fmt.Println(ips[0])

	resp, err := http.Get(fmt.Sprintf("http://%s:%s/json/version", chromeHost, chromePort))

	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		d, _ := ioutil.ReadAll(resp.Body)

		fmt.Println(string(d))
		return "", fmt.Errorf("Expected successful status code while fetching remote debugger, got %d", resp.StatusCode)
	}

	var jsonRes chromeHostJSONVersionResponseBody

	if err := json.NewDecoder(resp.Body).Decode(&jsonRes); err != nil {
		return "", err
	}

	webSocketDebuggerURL := jsonRes.WebSocketDebuggerURL

	if webSocketDebuggerURL == "" {
		return "", errors.New("Expected webSocketDebuggerUrl in /json/version response")
	}

	return webSocketDebuggerURL, nil
}

// IsAvailable check for PS5 availability
func (p CDiscountParser) IsAvailable() (bool, error) {
	chromeURL, err := p.getRemoteDebuggerURL()

	if err != nil {
		return false, err
	}

	// create allocator context for use with creating a browser context later
	allocatorContext, cancel := chromedp.NewRemoteAllocator(context.Background(), chromeURL)
	defer cancel()

	// create context
	ctxt, cancel := chromedp.NewContext(allocatorContext)
	defer cancel()

	// run task list
	var nodes []*cdp.Node
	if err := chromedp.Run(ctxt,
		chromedp.Navigate(p.URL),
		chromedp.WaitVisible(".prdBlocContainer"),
		chromedp.Nodes(`.crUl>.crItem`, &nodes, chromedp.ByQuery),
	); err != nil {
		return false, err
	}

	if len(nodes) != 1 {
		return true, nil
	}

	return false, nil
}