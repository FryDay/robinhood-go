package robinhood

import (
	"net/http"
)

const (
	baseURL = "https://api.robinhood.com"
)

// NewClient returns a new Robinhood client
func NewClient(config *Config) (*Client, error) {
	return &Client{
		config:     config,
		httpClient: http.DefaultClient,
	}, nil
}

// NewConfig returns a new config
func NewConfig(username, password string) *Config {
	return &Config{
		username: username,
		password: password,
	}
}
