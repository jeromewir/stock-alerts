package parsers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

// LeclercParser holds the informations needed to scrap fnac
type LeclercParser struct {
	Name string
	URL string
	shortURL string
}

type leclercProduct struct {
	Variants []struct {
		Offers []struct {
			BasePrice struct {
				Price struct {
					// price in cents
					Price float64
				}
			}
		}
	}
	Family struct {
		Code string
	}
}

type productSearchResponse struct {
	Count int
	Items []leclercProduct
}

type productSearchCategories struct {
	Code []string `json:"code"`
}

type productSearchRequestBody struct {
	Categories productSearchCategories `json:"categories"`
	Language string `json:"language"`
	Page int `json:"page"`
	Size int `json:"size"` // 90
	Sorts []string `json:"sorts"`
}

// NewLeclercParser returns a new instance of fnac parser
func NewLeclercParser() LeclercParser {
	return LeclercParser{
		Name: "Leclerc",
		URL: "https://www.e.leclerc/api/rest/live-api/product-search",
		shortURL: "https://cutt.ly/qjxix43",
	}
}

// GetName returns the display name
func (p LeclercParser) GetName() string {
	return p.Name
}

// GetURL returns the url of the scraper
func (p LeclercParser) GetURL() string {
	return p.URL
}

// GetShortURL returns the short URL used for display
func (p LeclercParser) GetShortURL() string {
	return p.shortURL
}

func (p LeclercParser) parseResponse(body io.ReadCloser) (*productSearchResponse, error) {
	var res productSearchResponse

	if err := json.NewDecoder(body).Decode(&res); err != nil {
		return nil, err
 	}

 	return &res, nil
}

func (p LeclercParser) getJSON(baseURL string) (*productSearchResponse, error) {
	requestBody, err := json.Marshal(productSearchRequestBody{
		Categories: productSearchCategories{Code: []string{"NAVIGATION_consoles-ps5"}},
		Language: "fr-FR",
		Page: 1,
		Size: 90,
		Sorts: []string{},
	})

	if err != nil {
		return nil, err
	}

	bodyBuffer := bytes.NewBuffer(requestBody)

	resp, err := http.Post(baseURL, "application/json", bodyBuffer)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return p.parseResponse(resp.Body)
}

// IsAvailable check for PS5 availability
func (p LeclercParser) IsAvailable() (bool, error) {
	productResponse, err := p.getJSON(p.URL)

	if err != nil {
		return false, err
	}

	if len(productResponse.Items) > 0 {
		return true, nil
	}

	return false, nil
}