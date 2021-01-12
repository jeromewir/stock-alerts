package parsers

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// BoulangerParser holds the informations needed to scrap fnac
type BoulangerParser struct {
	Name string
	URL string
	shortURL string
}

// NewBoulangerParser returns a new instance of fnac parser
func NewBoulangerParser() BoulangerParser {
	return BoulangerParser{
		Name: "Boulanger",
		URL: "https://www.boulanger.com/ref/1147567",
		shortURL: "https://cutt.ly/HjbTKpe",
	}
}

// GetName returns the name of the scraper
func (p BoulangerParser) GetName() string {
	return p.Name
}

// GetURL returns the url of the scraper
func (p BoulangerParser) GetURL() string {
	return p.URL
}

// GetShortURL returns the short URL used for display
func (p BoulangerParser) GetShortURL() string {
	return p.shortURL
}

func (p BoulangerParser) getPage(url string) (*goquery.Document, error) {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_0_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.101 Safari/537.36")

	return getDocument(req)
}


// IsAvailable check for PS5 availability
func (p BoulangerParser) IsAvailable() (bool, error) {
	page, err := p.getPage(p.URL)

	if err != nil {
		return false, err
	}

	unavailableDiv := page.Find(".infoAchat.delivery-purchase>.product-unavailable")

	if unavailableDiv.Length() == 1 {
		return false, nil
	}

	return true, nil
}