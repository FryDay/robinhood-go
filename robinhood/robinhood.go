package robinhood

import (
	"net/http"
)

const (
	baseURL = "https://api.robinhood.com"
)

// NewClient returns a new Robinhood client
func NewClient() (*Client, error) {
	return &Client{
		config:     &Config{},
		httpClient: http.DefaultClient,
	}, nil
}
