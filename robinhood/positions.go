package robinhood

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Positions gets positions data
func (c *Client) Positions() (*Positions, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/positions/", baseURL), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Token %s", c.config.Token))

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	positions := &Positions{}
	json.Unmarshal(content, positions)

	return positions, nil
}
