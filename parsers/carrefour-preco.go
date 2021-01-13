package parsers

import (
	"errors"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// CarrefourPrecoParser holds the informations needed to scrap fnac
type CarrefourPrecoParser struct {
	Name string
	URL string
	shortURL string
}

// NewCarrefourPrecoParser returns a new instance of fnac parser
func NewCarrefourPrecoParser() CarrefourPrecoParser {
	return CarrefourPrecoParser{
		Name: "Carrefour",
		URL: "https://reservation.carrefour.fr/",
		shortURL: "https://cutt.ly/AjzjXew",
	}
}

// GetName returns the name of the scraper
func (p CarrefourPrecoParser) GetName() string {
	return p.Name
}

// GetURL returns the url of the scraper
func (p CarrefourPrecoParser) GetURL() string {
	return p.URL
}

// GetShortURL returns the short URL used for display
func (p CarrefourPrecoParser) GetShortURL() string {
	return p.shortURL
}

func (p CarrefourPrecoParser) getPage(url string) (*goquery.Document, error) {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_0_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.101 Safari/537.36")

	return getDocument(req)
}

// IsAvailable check for PS5 availability
func (p CarrefourPrecoParser) IsAvailable() (bool, error) {
	document, err := p.getPage(p.URL)

  if err != nil {
    return false, err
	}

	sections := document.Find("section")

	if sections.Length() != 5 {
		return true, nil
	}

	previousCommandsSection := document.Find(".cm-section.cm-bg--gray-black.mts")

	if previousCommandsSection.Length() != 1 {
		return false, errors.New("Hum, looks like .cm-section.cm-bg--gray-black.mts is not present anymore")
	}

	previousCommandsRows := previousCommandsSection.Find("div.row")

	if previousCommandsRows.Length() != 1 {
		return true, nil
	}

	return false, nil
}