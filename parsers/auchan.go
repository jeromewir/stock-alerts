package parsers

import (
	"net/http"
)

// AuchanParser holds the informations needed to scrap fnac
type AuchanParser struct {
	Name string
	URL string
	shortURL string
}

// NewAuchanParser returns a new instance of fnac parser
func NewAuchanParser() AuchanParser {
	return AuchanParser{
		Name: "Auchan",
		URL: "https://www.auchan.fr/sony-console-ps5-edition-standard/p-c1315865?awc=7728_1609969871_7f94a075c1676c4473ea993241249d45&utm_medium=affiliation&utm_source=zanox&utm_campaign=generique&utm_content=0&utm_term=169249",
		shortURL: "https://cutt.ly/Pjzywnn",
	}
}

// GetName returns the name of the scraper
func (p AuchanParser) GetName() string {
	return p.Name
}

// GetURL returns the url of the scraper
func (p AuchanParser) GetURL() string {
	return p.URL
}

// GetShortURL returns the short URL used for display
func (p AuchanParser) GetShortURL() string {
	return p.shortURL
}

// IsAvailable check for PS5 availability
func (p AuchanParser) IsAvailable() (bool, error) {
	res, err := http.Get(p.URL)
  if err != nil {
    return false, err
	}

  defer res.Body.Close()

	if res.StatusCode == 404 {
		return false, nil
	}

	return true, nil
}