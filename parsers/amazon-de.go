package parsers

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// AmazonDEParser holds the informations needed to scrap fnac
type AmazonDEParser struct {
	Name string
	URL string
	shortURL string
}

// NewAmazonDEParser returns a new instance of fnac parser
func NewAmazonDEParser() AmazonDEParser {
	return AmazonDEParser{
		Name: "Amazon DE",
		URL: "https://www.amazon.de/Sony-Interactive-Entertainment-PlayStation-5/dp/B08H93ZRK9/ref=sr_1_2",
		shortURL: "https://cutt.ly/ajxjFQM",
	}
}

// GetName returns the name of the scraper
func (p AmazonDEParser) GetName() string {
	return p.Name
}

// GetURL returns the url of the scraper
func (p AmazonDEParser) GetURL() string {
	return p.URL
}

// GetShortURL returns the short URL used for display
func (p AmazonDEParser) GetShortURL() string {
	return p.shortURL
}

func (p AmazonDEParser) getPage(url string) (*goquery.Document, error) {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_0_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.101 Safari/537.36")

	return getDocument(req)
}

// IsAvailable check for PS5 availability
func (p AmazonDEParser) IsAvailable() (bool, error) {
	document, err := p.getPage(p.URL)

  if err != nil {
    return false, err
	}

	addToCart := document.Find("#add-to-cart-button")

	if addToCart.Length() == 0 {
		return false, nil
	}

	return true, nil
}