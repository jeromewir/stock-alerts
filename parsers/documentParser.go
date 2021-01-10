package parsers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/jeromewir/stockalerts/config"
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

func getRemoteDebuggerURL() (string, error) {
	chromeHost := config.GetChromeHost()
	chromePort := config.GetChromePort()

	if chromeHost != "localhost" {
		// https://github.com/skalfyfan/dockerized-puppeteer/commit/d31b2243ad1b22904bfa2f9f91ac075a3da0511a
		ips, _ := net.LookupIP(chromeHost)

		if len(ips) > 0 {
			// If we don't have ips or there is an error, we rely on the provided host
			chromeHost = ips[0].String()
		}
	}

	// Get webSocket URL
	resp, err := http.Get(fmt.Sprintf("http://%s:%s/json/version", chromeHost, chromePort))

	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		d, _ := ioutil.ReadAll(resp.Body)

		// Print the JSON body, is there something bad with the chrome headless ?
		log.Println(string(d))
		return "", fmt.Errorf("Expected successful status code while fetching remote debugger, got %d", resp.StatusCode)
	}

	var jsonRes chromeHostJSONVersionResponseBody

	if err := json.NewDecoder(resp.Body).Decode(&jsonRes); err != nil {
		return "", err
	}

	webSocketDebuggerURL := jsonRes.WebSocketDebuggerURL

	if webSocketDebuggerURL == "" {
		return "", errors.New("Expected webSocketDebuggerUrl in /json/version response")
	}

	return webSocketDebuggerURL, nil
}