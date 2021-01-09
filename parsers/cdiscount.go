package parsers

import (
	"net/http"

	"github.com/jeromewir/stockalerts/config"
)

// CDiscountParser holds the informations needed to scrap
type CDiscountParser struct {
	Name string
	URL string
	shortURL string
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

// IsAvailable check for PS5 availability
func (p CDiscountParser) IsAvailable() (bool, error) {
	req, err := http.NewRequest("GET", p.URL, nil)

	req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("referer", "https://www.cdiscount.com/jeux-pc-video-console/ps5/console-ps5/l-1035001.html")
	req.Header.Set("cache-control", "max-age=0")
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_0_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.101 Safari/537.36")
	req.Header.Set("accept-language", "en-US,en;q=0.9")
	req.Header.Set("sec-fetch-dest", "document")
	req.Header.Set("sec-fetch-mode", "navigate")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("sec-gpc", "1")
	 req.Header.Set("sec-fetch-user", "?1")
	 req.Header.Set("cookie", config.GetCDiscountCookie())

	if err != nil {
		return false, err
	}

	doc, err := getDocument(req)

	if err != nil {
		return false, err
	}

	selection := doc.Find(".crUl>.crItem")

	if selection.Length() != 1 {
		return true, nil
	}

	return false, nil
}