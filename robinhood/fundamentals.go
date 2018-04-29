package robinhood

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Fundamental gets fundamental data by ticker symbol
func (c *Client) Fundamental(tickerSymbol string) (*Fundamental, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/fundamentals/%s/", baseURL, tickerSymbol), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fundamental := &Fundamental{}
	json.Unmarshal(content, fundamental)

	return fundamental, nil
}

// Fundamentals gets fundamental data on multiple ticker symbols
func (c *Client) Fundamentals(tickerSymbols ...string) (*Fundamentals, error) {
	if len(tickerSymbols) > 10 {
		return nil, fmt.Errorf("Can only get 10 symbols at a time")
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/fundamentals/", baseURL), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")

	symbols := strings.Join(tickerSymbols, ",")
	q := req.URL.Query()
	q.Add("symbols", symbols)
	req.URL.RawQuery = q.Encode()

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fundamental := &Fundamentals{}
	json.Unmarshal(content, fundamental)

	return fundamental, nil
}
