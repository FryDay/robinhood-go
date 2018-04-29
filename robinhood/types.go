package robinhood

import "net/http"

type (
	// Client represents a Robinhood API client
	Client struct {
		config     *Config
		httpClient *http.Client
	}

	// Config holds our clients information
	Config struct {
		username string
		password string
		Token    string `json:"token"`
	}

	// Fundamental data
	Fundamental struct {
		Open              *float64 `json:"open,string"`
		High              *float64 `json:"high,string"`
		Low               *float64 `json:"low,string"`
		Volume            *float64 `json:"volume,string"`
		AverageVolume     *float64 `json:"average_volume,string"`
		High52Weeks       *float64 `json:"high_52_weeks,string"`
		Low52Weeks        *float64 `json:"low_52_weeks,string"`
		MarketCap         *float64 `json:"market_cap,string"`
		DividendYield     *float64 `json:"dividend_yield,string"`
		PERatio           *float64 `json:"pe_ratio,string"`
		SharesOutstanding *float64 `json:"shares_outstanding,string"`
		Description       *string  `json:"description"`
		Instrument        *string  `json:"instrument"`
		CEO               *string  `json:"ceo"`
		HeadquartersCity  *string  `json:"headquarters_city"`
		HeadquartersState *string  `json:"headquarters_state"`
		Sector            *string  `json:"sector"`
		NumEmployees      *int64   `json:"num_employees"`
		YearFounded       *int64   `json:"year_founded"`
	}

	// Fundamentals data
	Fundamentals struct {
		Fundamentals []*Fundamental `json:"results"`
	}
)
