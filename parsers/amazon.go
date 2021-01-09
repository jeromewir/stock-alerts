package parsers

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// AmazonParser holds the informations needed to scrap fnac
type AmazonParser struct {
	Name string
	URL string
	shortURL string
}

// NewAmazonParser returns a new instance of fnac parser
func NewAmazonParser() AmazonParser {
	return AmazonParser{
		Name: "Amazon",
		URL: "https://www.amazon.fr/dp/B08H93ZRK9/ref=twister_B08HJZNMF3?_encoding=UTF8&th=1",
		shortURL: "https://cutt.ly/qjzsNUR",
	}
}

// GetName returns the name of the scraper
func (p AmazonParser) GetName() string {
	return p.Name
}

// GetURL returns the url of the scraper
func (p AmazonParser) GetURL() string {
	return p.URL
}

// GetShortURL returns the short URL used for display
func (p AmazonParser) GetShortURL() string {
	return p.shortURL
}

func (p AmazonParser) getPage(url string) (*goquery.Document, error) {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_0_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.101 Safari/537.36")

	return getDocument(req)
}

// IsAvailable check for PS5 availability
func (p AmazonParser) IsAvailable() (bool, error) {
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