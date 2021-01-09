package parsers

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/jeromewir/stockalerts/config"
)

// FnacParser holds the informations needed to scrap fnac
type FnacParser struct {
	Name string
	URL string
	shortURL string
}

// NewFnacParser returns a new instance of fnac parser
func NewFnacParser() FnacParser {
	return FnacParser{
		Name: "Fnac",
		URL: "https://www.fnac.com/Console-Sony-PS5-Edition-Standard/a14119956/w-4",
		shortURL: "https://cutt.ly/WjzyiPT",
	}
}

// GetName returns the display name
func (p FnacParser) GetName() string {
	return p.Name
}

// GetURL returns the url of the scraper
func (p FnacParser) GetURL() string {
	return p.URL
}

// GetShortURL returns the short URL used for display
func (p FnacParser) GetShortURL() string {
	return p.shortURL
}

func (p FnacParser) getPage(url string) (*goquery.Document, error) {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("authority", "www.fnac.com")
	req.Header.Set("cache-control", "max-age=0")
	req.Header.Set("upgrade-insecure-requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_0_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.101 Safari/537.36")
	req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("sec-gpc", "1")
	req.Header.Set("sec-fetch-site", "none")
	req.Header.Set("sec-fetch-mode", "navigate")
	req.Header.Set("sec-fetch-user", "?1")
	req.Header.Set("sec-fetch-dest", "document")
	req.Header.Set("accept-language", " en-US,en;q=0.9")
	req.Header.Set("cookie", config.GetFnacCookie())

	return getDocument(req)
}

// IsAvailable check for PS5 availability
func (p FnacParser) IsAvailable() (bool, error) {
	document, err := p.getPage(p.URL)

	if err != nil {
		return false, err
	}

	selection := document.Find(".f-buyBox-infos>.f-buyBox-availabilityStatus-unavailable")

	if selection.Length() != 2 {
		return true, nil
	}

	return false, nil
}