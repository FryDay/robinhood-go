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

// User gets user info
func (c *Client) User() (*User, error) {
	if len(c.config.Token) == 0 {
		return nil, fmt.Errorf("Not logged in")
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/user/", baseURL), nil)
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

	user := &User{}
	json.Unmarshal(content, user)

	return user, nil
}

// UserBasicInfo gets basic user info
func (c *Client) UserBasicInfo() (*UserBasicInfo, error) {
	if len(c.config.Token) == 0 {
		return nil, fmt.Errorf("Not logged in")
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/user/basic_info/", baseURL), nil)
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

	userBasicInfo := &UserBasicInfo{}
	json.Unmarshal(content, userBasicInfo)

	return userBasicInfo, nil
}

// UserAdditionalInfo gets additional user info
func (c *Client) UserAdditionalInfo() (*UserAdditionalInfo, error) {
	if len(c.config.Token) == 0 {
		return nil, fmt.Errorf("Not logged in")
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/user/additional_info/", baseURL), nil)
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

	userAdditionalInfo := &UserAdditionalInfo{}
	json.Unmarshal(content, userAdditionalInfo)

	return userAdditionalInfo, nil
}

// UserEmployment gets user employment info
func (c *Client) UserEmployment() (*UserEmployment, error) {
	if len(c.config.Token) == 0 {
		return nil, fmt.Errorf("Not logged in")
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/user/employment/", baseURL), nil)
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

	userEmployment := &UserEmployment{}
	json.Unmarshal(content, userEmployment)

	return userEmployment, nil
}

// UserInvestmentProfile gets user investment profile
func (c *Client) UserInvestmentProfile() (*UserInvestmentProfile, error) {
	if len(c.config.Token) == 0 {
		return nil, fmt.Errorf("Not logged in")
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/user/investment_profile/", baseURL), nil)
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

	userInvestmentProfile := &UserInvestmentProfile{}
	json.Unmarshal(content, userInvestmentProfile)

	return userInvestmentProfile, nil
}
