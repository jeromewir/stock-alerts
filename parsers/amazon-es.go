package parsers

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// AmazonESParser holds the informations needed to scrap fnac
type AmazonESParser struct {
	Name string
	URL string
	shortURL string
}

// NewAmazonESParser returns a new instance of fnac parser
func NewAmazonESParser() AmazonESParser {
	return AmazonESParser{
		Name: "Amazon ES",
		URL: "https://www.amazon.es/dp/B08KKJ37F7/ref=sxts_spkl_2_0_19702538-81b7-4e26-8a68-a475cc66898a",
		shortURL: "https://cutt.ly/vjxjIV9",
	}
}

// GetName returns the name of the scraper
func (p AmazonESParser) GetName() string {
	return p.Name
}

// GetURL returns the url of the scraper
func (p AmazonESParser) GetURL() string {
	return p.URL
}

// GetShortURL returns the short URL used for display
func (p AmazonESParser) GetShortURL() string {
	return p.shortURL
}

func (p AmazonESParser) getPage(url string) (*goquery.Document, error) {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_0_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.101 Safari/537.36")

	return getDocument(req)
}

// IsAvailable check for PS5 availability
func (p AmazonESParser) IsAvailable() (bool, error) {
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