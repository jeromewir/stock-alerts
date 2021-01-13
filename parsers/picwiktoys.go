package parsers

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/google/uuid"
)

// PicwictoysParser holds the informations needed to scrap fnac
type PicwictoysParser struct {
	Name string
	URL string
	shortURL string
}

type picwictoysSearchBody struct {
	ModuleVersion string `json:"moduleVersion"`
	SessionID     string `json:"sessionId"`
	PredSessionID string `json:"predSessionId"`
	Region        string `json:"region"`
	PageID        int    `json:"pageId"`
	Parameters    struct {
		Query          string `json:"query"`
		ResultsPerPage int    `json:"resultsPerPage"`
	} `json:"parameters"`
}

type picwictoysSearchBodyResult struct {
	Status    string  `json:"status"`
	TimeInMs  float64 `json:"timeInMs"`
	SessionID string  `json:"sessionId"`
	Region    string  `json:"region"`
	PageID    int     `json:"pageId"`
	PageName  string  `json:"pageName"`
	Blocks    struct {
		Ads             []interface{} `json:"ads"`
		Recommendations []interface{} `json:"recommendations"`
		Searches        []struct {
			BlockID    int           `json:"blockId"`
			PageName   string        `json:"pageName"`
			BlockName  string        `json:"blockName"`
			BlockType  string        `json:"blockType"`
			DidYouMean []interface{} `json:"didYouMean"`
			Stats      struct {
				TotalResults     int     `json:"totalResults"`
				ComputationInMs  float64 `json:"computationInMs"`
				ResultsInPage    int     `json:"resultsInPage"`
				ResultsPerPage   int     `json:"resultsPerPage"`
				CurrentPage      int     `json:"currentPage"`
				SearchID         int     `json:"searchId"`
				GalaxyID         string  `json:"galaxyId"`
				SortingCode      string  `json:"sortingCode"`
				TotalPages       int     `json:"totalPages"`
				RuleTemplate     string  `json:"ruleTemplate"`
				Query            string  `json:"query"`
				UpdatedQuery     string  `json:"updatedQuery"`
				ScenarioTemplate string  `json:"scenarioTemplate"`
				NbSlotsX         int     `json:"nbSlotsX"`
				NbSlotsY         int     `json:"nbSlotsY"`
			} `json:"stats"`
			Facets []struct {
				AttributeName  string `json:"attributeName"`
				AttributeLabel string `json:"attributeLabel"`
				AttributeType  string `json:"attributeType"`
				AttributeUnit  string `json:"attributeUnit"`
				MultiSelect    bool   `json:"multiSelect"`
				Position       int    `json:"position"`
				Options        []struct {
					RefiningID string `json:"refiningId"`
					NoFollow   bool   `json:"noFollow"`
					NbResults  int    `json:"nbResults"`
					Selected   bool   `json:"selected"`
					DebugText  string `json:"debugText"`
					Value      string `json:"value"`
					Label      string `json:"label"`
				} `json:"options"`
			} `json:"facets"`
			Backtracks []interface{} `json:"backtracks"`
			Pages      []struct {
				RefiningID string `json:"refiningId"`
				NoFollow   bool   `json:"noFollow"`
				PageNumber int    `json:"pageNumber"`
				Label      string `json:"label"`
			} `json:"pages"`
			Sortings []struct {
				RefiningID  string `json:"refiningId"`
				NoFollow    bool   `json:"noFollow"`
				SortingCode string `json:"sortingCode"`
				Selected    bool   `json:"selected"`
				Label       string `json:"label"`
			} `json:"sortings"`
			Slots []struct {
				Position int `json:"position"`
				Item     struct {
					Sku           string  `json:"sku"`
					Price         float64 `json:"price"`
					ItemType      string  `json:"itemType"`
					TrackingCode  string  `json:"trackingCode"`
					AttributeInfo []struct {
						AttributeName  string `json:"attributeName"`
						AttributeLabel string `json:"attributeLabel"`
						AttributeType  string `json:"attributeType"`
						Vals           []struct {
							Value string `json:"value"`
							Label string `json:"label"`
						} `json:"vals"`
					} `json:"attributeInfo"`
					DisplayInfo struct {
					} `json:"displayInfo"`
				} `json:"item"`
				Template struct {
					PosX  int `json:"posX"`
					PosY  int `json:"posY"`
					SizeX int `json:"sizeX"`
					SizeY int `json:"sizeY"`
				} `json:"template"`
			} `json:"slots"`
			NbSlots int `json:"nbSlots"`
		} `json:"searches"`
	} `json:"blocks"`
	NbItems int `json:"nbItems"`
}

// NewPicwictoysParser returns a new instance of fnac parser
func NewPicwictoysParser() PicwictoysParser {
	return PicwictoysParser{
		Name: "Picwiktoys",
		URL: "https://pictoys-prod.prediggo.io/ptoys-prod-JS-WLzZEeEQpU84gJ7yHCAv/3.0/simplePageContent",
		shortURL: "https://cutt.ly/PjmADEk",
	}
}

// GetName returns the display name
func (p PicwictoysParser) GetName() string {
	return p.Name
}

// GetURL returns the url of the scraper
func (p PicwictoysParser) GetURL() string {
	return p.URL
}

// GetShortURL returns the short URL used for display
func (p PicwictoysParser) GetShortURL() string {
	return p.shortURL
}

func (p PicwictoysParser) parseResponse(body io.ReadCloser) (*picwictoysSearchBodyResult, error) {
	var res picwictoysSearchBodyResult

	if err := json.NewDecoder(body).Decode(&res); err != nil {
		return nil, err
 	}

 	return &res, nil
}

func (p PicwictoysParser) getJSON(baseURL string) (*picwictoysSearchBodyResult, error) {
	sessionID := uuid.New().String()
	prevSessionID := uuid.New().String()

	requestBody, err := json.Marshal(picwictoysSearchBody{
		ModuleVersion: "production",
		SessionID: sessionID,
		PredSessionID: prevSessionID,
		Region: "fr_FR",
		PageID: 1,
		Parameters: struct {
			Query          string `json:"query"`
			ResultsPerPage int    `json:"resultsPerPage"`
		}{ Query: "ps5", ResultsPerPage: 18 },
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
func (p PicwictoysParser) IsAvailable() (bool, error) {
	productResponse, err := p.getJSON(p.URL)

	if err != nil {
		return false, err
	}

	if productResponse != nil && len(productResponse.Blocks.Searches) > 0 {
		for _, slot := range productResponse.Blocks.Searches[0].Slots {
			if slot.Item.Price == 499.99 {
				return true, nil
			}
		}
	} else {
		return false, errors.New("Unexpected productResponse")
	}

	return false, nil
}