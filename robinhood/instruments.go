package robinhood

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// InstrumentBySymbol gets instrument data by ticker symbol
func (c *Client) InstrumentBySymbol(tickerSymbol string) (*Instrument, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s?symbol=%s", baseURL, "/instruments/", tickerSymbol), nil)
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
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s%s/", baseURL, "/instruments/", id), nil)
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
