package parsers

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// CarrefourParser holds the informations needed to scrap fnac
type CarrefourParser struct {
	Name string
	URL string
	shortURL string
}

// NewCarrefourParser returns a new instance of fnac parser
func NewCarrefourParser() CarrefourParser {
	return CarrefourParser{
		Name: "Carrefour",
		URL: "https://jouetsdenoel.carrefour.fr/produit/jeux-video-et-multimedia/ps5-cons-ps5-standard",
		shortURL: "https://cutt.ly/njmY4bf",
	}
}

// GetName returns the name of the scraper
func (p CarrefourParser) GetName() string {
	return p.Name
}

// GetURL returns the url of the scraper
func (p CarrefourParser) GetURL() string {
	return p.URL
}

// GetShortURL returns the short URL used for display
func (p CarrefourParser) GetShortURL() string {
	return p.shortURL
}

func (p CarrefourParser) getPage(url string) (*goquery.Document, error) {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_0_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.101 Safari/537.36")

	return getDocument(req)
}

// IsAvailable check for PS5 availability
func (p CarrefourParser) IsAvailable() (bool, error) {
	document, err := p.getPage(p.URL)

  if err != nil {
    return false, err
	}

	addToCartBtn := document.Find(".cm-product.js-product-detail .cm-pastille-container>a.product-add")

	if addToCartBtn.Length() > 0 {
		return true, nil
	}

	return false, nil
}