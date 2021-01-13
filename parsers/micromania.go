package parsers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// MicromaniaParser holds the informations needed to scrap fnac
type MicromaniaParser struct {
	Name string
	URL string
	shortURL string
}

// NewMicromaniaParser returns a new instance of fnac parser
func NewMicromaniaParser() MicromaniaParser {
	return MicromaniaParser{
		Name: "Micromania",
		URL: "https://www.micromania.fr/playstation-5-105642.html",
		shortURL: "https://cutt.ly/djmTpks",
	}
}

// GetName returns the name of the scraper
func (p MicromaniaParser) GetName() string {
	return p.Name
}

// GetURL returns the url of the scraper
func (p MicromaniaParser) GetURL() string {
	return p.URL
}

// GetShortURL returns the short URL used for display
func (p MicromaniaParser) GetShortURL() string {
	return p.shortURL
}

func (p MicromaniaParser) getPage(url string) (*goquery.Document, error) {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_0_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.101 Safari/537.36")
	req.Header.Set("accept-language", "en-US,en;q=0.9")

	return getDocument(req)
}


// IsAvailable check for PS5 availability
func (p MicromaniaParser) IsAvailable() (bool, error) {
	page, err := p.getPage(p.URL)

	if err != nil {
		return false, err
	}

	saleBlock := page.Find(".add-to-cart-container")

	if saleBlock.Length() != 1 {
		return false, errors.New("Expected .add-to-cart-container to be defined")
	}

	classValue, exists := saleBlock.Attr("class")

	if !exists {
		return false, errors.New("Expected to find .saleBlock.d-none div")
	}

	if strings.Contains(classValue, "d-none") {
		return false, nil
	}

	return true, nil
}