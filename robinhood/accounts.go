package robinhood

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Account gets account data
func (c *Client) Account() (*Account, error) {
	if len(c.config.Token) == 0 {
		return nil, fmt.Errorf("Not logged in")
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/accounts/", baseURL), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Token %s", c.config.Token))
	fmt.Println(c.config.Token)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	accounts := &Accounts{}
	json.Unmarshal(content, accounts)
	if len(accounts.Accounts) == 0 {
		return nil, nil
	}

	return accounts.Accounts[0], nil
}
