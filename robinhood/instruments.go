package robinhood

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// InstrumentBySymbol gets instrument data by ticker symbol
func (c *Client) InstrumentBySymbol(tickerSymbol string) (*Instrument, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/instruments/?symbol=%s", baseURL, tickerSymbol), nil)
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

	instruments := &Instruments{}
	json.Unmarshal(content, instruments)
	if len(instruments.Instruments) == 0 {
		return nil, nil
	}

	return instruments.Instruments[0], nil
}

// InstrumentByID gets instrument data by ID
func (c *Client) InstrumentByID(id string) (*Instrument, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/instruments/%s/", baseURL, id), nil)
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

	instrument := &Instrument{}
	json.Unmarshal(content, instrument)

	return instrument, nil
}

//TODO: Search by keyword

// Splits gets splits by ID
func (c *Client) Splits(id string) (*Splits, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/instruments/%s/splits/", baseURL, id), nil)
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

	splits := &Splits{}
	json.Unmarshal(content, splits)

	return splits, nil
}

// Split gets a single split
func (c *Client) Split(id, splitID string) (*Split, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/instruments/%s/splits/%s/", baseURL, id, splitID), nil)
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

	split := &Split{}
	json.Unmarshal(content, split)

	return split, nil
}
