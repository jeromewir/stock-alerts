package parsers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// LDLCParser holds the informations needed to scrap fnac
type LDLCParser struct {
	Name string
	URL string
	shortURL string
}

// NewLDLCParser returns a new instance of fnac parser
func NewLDLCParser() LDLCParser {
	return LDLCParser{
		Name: "LDLC",
		URL: "https://www.ldlc.com/fiche/PB00386358.html",
		shortURL: "https://cutt.ly/ojbYPOS",
	}
}

// GetName returns the name of the scraper
func (p LDLCParser) GetName() string {
	return p.Name
}

// GetURL returns the url of the scraper
func (p LDLCParser) GetURL() string {
	return p.URL
}

// GetShortURL returns the short URL used for display
func (p LDLCParser) GetShortURL() string {
	return p.shortURL
}

func (p LDLCParser) getPage(url string) (*goquery.Document, error) {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_0_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.101 Safari/537.36")
	req.Header.Set("accept-language", "en-US,en;q=0.9")

	return getDocument(req)
}


// IsAvailable check for PS5 availability
func (p LDLCParser) IsAvailable() (bool, error) {
	page, err := p.getPage(p.URL)

	if err != nil {
		return false, err
	}

	saleBlock := page.Find(".saleBlock")

	if saleBlock.Length() != 1 {
		return false, errors.New("Expected .saleBlock to be defined")
	}

	classValue, exists := saleBlock.Attr("class")

	if !exists {
		return false, errors.New("Expected to find .saleBlock.invisible div")
	}

	if strings.Contains(classValue, "invisible") {
		return false, nil
	}

	return true, nil
}