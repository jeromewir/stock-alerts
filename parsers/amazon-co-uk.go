package parsers

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// AmazonCOUKParser holds the informations needed to scrap fnac
type AmazonCOUKParser struct {
	Name string
	URL string
	shortURL string
}

// NewAmazonCOUKParser returns a new instance of fnac parser
func NewAmazonCOUKParser() AmazonCOUKParser {
	return AmazonCOUKParser{
		Name: "Amazon UK",
		URL: "https://www.amazon.co.uk/PlayStation-9395003-5-Console/dp/B08H95Y452/ref=sr_1_1?dchild=1&keywords=ps5&qid=1610473820&sr=8-1",
		shortURL: "https://cutt.ly/ijbThkH",
	}
}

// GetName returns the name of the scraper
func (p AmazonCOUKParser) GetName() string {
	return p.Name
}

// GetURL returns the url of the scraper
func (p AmazonCOUKParser) GetURL() string {
	return p.URL
}

// GetShortURL returns the short URL used for display
func (p AmazonCOUKParser) GetShortURL() string {
	return p.shortURL
}

func (p AmazonCOUKParser) getPage(url string) (*goquery.Document, error) {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_0_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.101 Safari/537.36")

	return getDocument(req)
}

// IsAvailable check for PS5 availability
func (p AmazonCOUKParser) IsAvailable() (bool, error) {
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