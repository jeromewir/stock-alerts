package parsers

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func getDocument(req *http.Request) (*goquery.Document, error) {
	client := &http.Client{}

	res, err := client.Do(req)

  if err != nil {
    return nil, err
  }

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
  }

  // Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)

  if err != nil {
		return nil, err
	}
	
	return doc, err
}