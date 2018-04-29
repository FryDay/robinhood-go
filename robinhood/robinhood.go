package robinhood

import (
	"net/http"
)

const (
	baseURL = "https://api.robinhood.com"
)

// Client represents a Robinhood API client
type Client struct {
	config     *Config
	httpClient *http.Client
}

// Config holds our clients information
type Config struct {
	username string
	password string
	Token    string `json:"token"`
}

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
