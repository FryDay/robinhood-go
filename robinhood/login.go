package robinhood

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Login sets token based on a username and password
// Endpoint: POST /api-token-auth/
func (c *Client) Login(username, password string) error {
	v := url.Values{}
	v.Set("username", username)
	v.Set("password", password)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", baseURL, "/api-token-auth/"), strings.NewReader(v.Encode()))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	json.Unmarshal(content, c.config)
	if len(c.config.Token) == 0 {
		return fmt.Errorf("Could not login")
	}

	return nil
}
