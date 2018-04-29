package robinhood

import (
	"net/http"
	"time"
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

	// Account data
	Account struct {
		Deactivated                *bool           `json:"deactivated"`
		UpdatedAt                  *time.Time      `json:"updated_at"`
		MarginBalances             *MarginBalances `json:"margin_balances"`
		Portfolio                  *string         `json:"portfolio"`
		CashBalances               *CashBalances   `json:"cash_balances"`
		WithdrawalHalted           *bool           `json:"withdrawal_halted"`
		CashAvailableForWithdrawal *float64        `json:"cash_available_for_withdrawal,string"`
		Type                       *string         `json:"type"`
		SMA                        *float64        `json:"sma,string"`
		SweepEnabled               *bool           `json:"sweep_enabled"`
		DepositHalted              *bool           `json:"deposit_halted"`
		BuyingPower                *float64        `json:"buying_power,string"`
		User                       *string         `json:"user"`
		MaxAchEarlyAccessAmount    *float64        `json:"max_ach_early_access_amount,string"`
		//TODO: Option Level?
		//TODO: Instant Eligibility?
		CashHeldForOrders         *float64   `json:"cash_held_for_orders,string"`
		OnlyPositionClosingTrades *bool      `json:"only_position_closing_trades"`
		URL                       *string    `json:"url"`
		Positions                 *string    `json:"positions"`
		CreatedAt                 *time.Time `json:"created_at"`
		Cash                      *float64   `json:"cash,string"`
		SMAHeldForOrders          *float64   `json:"sma_held_for_orders,string"`
		UnsettledDebit            *float64   `json:"unsettled_debit,string"`
		AccountNumber             *string    `json:"account_number"`
		UnclearedDeposits         *float64   `json:"uncleared_deposits,string"`
		UnsettledFunds            *float64   `json:"unsettled_funds,string"`
	}

	// Accounts data
	Accounts struct {
		Previous *string    `json:"previous"`
		Accounts []*Account `json:"results"`
		Next     *string    `json:"next"`
	}

	// MarginBalances data
	MarginBalances struct {
		UpdatedAt                         *time.Time `json:"updated_at"`
		GoldEquityRequirement             *float64   `json:"gold_equity_requirement,string"`
		OutstandingInterest               *float64   `json:"outstanding_interest,string"`
		CashHeldForOptionsCollateral      *float64   `json:"cash_held_for_options_collateral,string"`
		OvernightRatio                    *float64   `json:"overnight_ratio,string"`
		DayTradeBuyingPower               *float64   `json:"day_trade_buying_power,string"`
		CashAvailableForWithdrawal        *float64   `json:"cash_available_for_withdrawal,string"`
		SMA                               *float64   `json:"sma,string"`
		MarkedPatternDayTraderDate        *time.Time `json:"marked_pattern_day_trader_date"`
		UnallocatedMarginCash             *float64   `json:"unallocated_margin_cash,string"`
		UnclearedDeposits                 *float64   `json:"uncleared_deposits,string"`
		OvernightBuyingPowerHeldForOrders *float64   `json:"overnight_buying_power_held_for_orders,string"`
		DayTradeRatio                     *float64   `json:"day_trade_ratio,string"`
		CashHeldForOrders                 *float64   `json:"cash_held_for_orders,string"`
		UnsettledDebit                    *float64   `json:"unsettled_debit,string"`
		CreatedAt                         *time.Time `json:"created_at"`
		Cash                              *float64   `json:"cash,string"`
		StartOfDayOvernightBuyingPower    *float64   `json:"start_of_day_overnight_buying_power,string"`
		MarginLimit                       *float64   `json:"margin_limit,string"`
		OvernightBuyingPower              *float64   `json:"overnight_buying_power,string"`
		StartOfDayDTBP                    *float64   `json:"start_of_day_dtbp,string"`
		UnsettledFunds                    *float64   `json:"unsettled_funds,string"`
		DayTradeBuyingPowerHeldForOrders  *float64   `json:"day_trade_buying_power_held_for_orders,string"`
	}

	// CashBalances data
	CashBalances struct {
		CashHeldForOrders          *float64   `json:"cash_held_for_orders,string"`
		CreatedAt                  *time.Time `json:"created_at"`
		Cash                       *float64   `json:"cash,string"`
		BuyingPower                *float64   `json:"buying_power,string"`
		UpdatedAt                  *time.Time `json:"updated_at"`
		CashAvailableForWithdrawal *float64   `json:"cash_available_for_withdrawl,string"`
		UnclearedDeposits          *float64   `json:"uncleared_deposits,string"`
		UnsettledFunds             *float64   `json:"unsettled_funds,string"`
	}

	// User data
	User struct {
		Username          *string    `json:"username"`
		FirstName         *string    `json:"first_name"`
		LastName          *string    `json:"last_name"`
		IDInfo            *string    `json:"id_info"`
		URL               *string    `json:"url"`
		EmailVerified     *bool      `json:"email_verified"`
		CreatedAt         *time.Time `json:"created_at"`
		BasicInfo         *string    `json:"basic_info"`
		Email             *string    `json:"email"`
		InvestmentProfile *string    `json:"investment_profile"`
		ID                *string    `json:"id"`
		InternationalInfo *string    `json:"international_info"`
		Employment        *string    `json:"employment"`
		AdditionalInfo    *string    `json:"additional_info"`
	}
)
