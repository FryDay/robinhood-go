package robinhood

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Fundamental gets fundamental data by ticket symbol
func (c *Client) Fundamental(tickerSymbol string) (*Fundamental, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s%s/", baseURL, "/fundamentals/", tickerSymbol), nil)
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
