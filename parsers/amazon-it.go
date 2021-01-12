package parsers

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// AmazonITParser holds the informations needed to scrap fnac
type AmazonITParser struct {
	Name string
	URL string
	shortURL string
}

// NewAmazonITParser returns a new instance of fnac parser
func NewAmazonITParser() AmazonITParser {
	return AmazonITParser{
		Name: "Amazon IT",
		URL: "https://www.amazon.it/dp/B08KKJ37F7/",
		shortURL: "https://cutt.ly/IjbTEqP",
	}
}

// GetName returns the name of the scraper
func (p AmazonITParser) GetName() string {
	return p.Name
}

// GetURL returns the url of the scraper
func (p AmazonITParser) GetURL() string {
	return p.URL
}

// GetShortURL returns the short URL used for display
func (p AmazonITParser) GetShortURL() string {
	return p.shortURL
}

func (p AmazonITParser) getPage(url string) (*goquery.Document, error) {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_0_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.101 Safari/537.36")

	return getDocument(req)
}

// IsAvailable check for PS5 availability
func (p AmazonITParser) IsAvailable() (bool, error) {
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