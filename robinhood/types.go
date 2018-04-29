package robinhood

import (
	"net/http"
)

type (
	// Client represents a Robinhood API client
	Client struct {
		config     *Config
		httpClient *http.Client
	}

	// Config holds our clients information
	Config struct {
		Token string `json:"token"`
	}

	// Fundamental data
	Fundamental struct {
		Open                *float64 `json:"open,string"`
		High                *float64 `json:"high,string"`
		Low                 *float64 `json:"low,string"`
		Volume              *float64 `json:"volume,string"`
		AverageVolume2Weeks *float64 `json:"average_volume_2_weeks,string"`
		AverageVolume       *float64 `json:"average_volume,string"`
		High52Weeks         *float64 `json:"high_52_weeks,string"`
		DividendYield       *float64 `json:"dividend_yield,string"`
		Low52Weeks          *float64 `json:"low_52_weeks,string"`
		MarketCap           *float64 `json:"market_cap,string"`
		PERatio             *float64 `json:"pe_ratio,string"`
		SharesOutstanding   *float64 `json:"shares_outstanding,string"`
		Description         *string  `json:"description"`
		Instrument          *string  `json:"instrument"`
		CEO                 *string  `json:"ceo"`
		HeadquartersCity    *string  `json:"headquarters_city"`
		HeadquartersState   *string  `json:"headquarters_state"`
		Sector              *string  `json:"sector"`
		NumEmployees        *int64   `json:"num_employees"`
		YearFounded         *int64   `json:"year_founded"`
	}

	// Fundamentals data
	Fundamentals struct {
		Fundamentals []*Fundamental `json:"results"`
	}

	// Instrument data
	Instrument struct {
		MinTickSize        *float64 `json:"min_tick_size,string"`
		Type               *string  `json:"type"`
		Splits             *string  `json:"splits"`
		MarginInitialRatio *float64 `json:"margin_initial_ratio,string"`
		URL                *string  `json:"url"`
		Quote              *string  `json:"quote"`
		Tradability        *string  `json:"tradability"`
		BloombergUnique    *string  `json:"bloomberg_unique"`
		ListDate           *date    `json:"list_date"`
		Name               *string  `json:"name"`
		Symbol             *string  `json:"symbol"`
		Fundamentals       *string  `json:"fundamentals"`
		State              *string  `json:"state"`
		Country            *string  `json:"country"`
		DayTradeRatio      *string  `json:"day_trade_ratio"`
		Tradeable          *bool    `json:"tradeable"`
		MaintenanceRatio   *float64 `json:"maintenance_ratio,string"`
		ID                 *string  `json:"id"`
		Market             *string  `json:"market"`
		SimpleName         *string  `json:"simple_name"`
	}

	// Instruments data
	Instruments struct {
		Previous    *string       `json:"previous"`
		Instruments []*Instrument `json:"results"`
		Next        *string       `json:"next"`
	}

	// Split data
	Split struct {
		URL           *string  `json:"url"`
		Instrument    *string  `json:"instrument"`
		ExecutionDate *date    `json:"execution_date"`
		Divisor       *float64 `json:"divisor,string"`
		Multiplier    *float64 `json:"multiplier,string"`
	}

	// Splits data
	Splits struct {
		Previous *string  `json:"previous"`
		Splits   []*Split `json:"results"`
		Next     *string  `json:"next"`
	}
)
