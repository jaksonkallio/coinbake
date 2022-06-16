package service

import (
	"fmt"

	"github.com/neiltcox/coinbake/database"
	"gorm.io/gorm"
)

type ExchangeIdentifier string

const (
	ExchangeIdentifierMocked      ExchangeIdentifier = "mock"
	ExchangeIdentifierKraken      ExchangeIdentifier = "kraken"
	ExchangeIdentifierCoinbasePro ExchangeIdentifier = "coinbasepro"
	ExchangeIdentifierBinance     ExchangeIdentifier = "binance"
)

var exchanges map[ExchangeIdentifier]Exchange = make(map[ExchangeIdentifier]Exchange)

func init() {
	exchanges[ExchangeIdentifierMocked] = &ExchangeMocked{}
	exchanges[ExchangeIdentifierKraken] = &ExchangeKraken{}
}

// Represents a configured connection to an exchange.
type Portfolio struct {
	gorm.Model
	ApiKey    string
	ApiSecret string
	ExchangeIdentifier
	UserID int
	User   User
}

type SupportedAsset struct {
	Asset Asset
}

// An interface representing a generic exchange.
type Exchange interface {
	CreateOrder(*Portfolio, string, float32) (CreatedOrder, error)
	Holdings(*Portfolio) ([]Holding, error)
	SupportedAssets(*Portfolio) (map[string]bool, error)
}

type MockSupportedAssets struct {
}

type CreatedOrder struct {
	OrderIdentifier string
}

type Holding struct {
	Asset   string
	Balance float32
}

type ExchangeKraken struct {
}

type ExchangeMocked struct {
	MockSupportedAssets
}

func (mockSupportedAssets *MockSupportedAssets) SupportedAssets(exchangeConnection *Portfolio) (map[string]bool, error) {
	return map[string]bool{
		"BTC": true,
		"ETH": true,
		"XMR": true,
	}, nil
}

func (exchangeKraken *ExchangeKraken) SupportedAssets(exchangeConnection *Portfolio) (map[string]bool, error) {
	// TODO: implement
	return map[string]bool{}, nil
}

func FindPortfoliosByUserId(userId uint) []Portfolio {
	portfolios := []Portfolio{}
	database.Handle().Where("user_id = ?", userId).Find(&portfolios)
	return portfolios
}

// Gets the Exchange object for a given Exchange Connection, which is where the API call logic is.
func (exchangeConnection *Portfolio) Exchange() (Exchange, error) {
	exchange, exists := exchanges[exchangeConnection.ExchangeIdentifier]
	if !exists {
		return nil, fmt.Errorf("exchange %q is not implemented", exchangeConnection.ExchangeIdentifier)
	}

	return exchange, nil
}

func (exchangeMocked *ExchangeMocked) CreateOrder(exchangeConnection *Portfolio, asset string, amount float32) (CreatedOrder, error) {
	return CreatedOrder{
		OrderIdentifier: "123456",
	}, nil
}

func (exchangeMocked *ExchangeMocked) Holdings(exchangeConnection *Portfolio) ([]Holding, error) {
	return []Holding{
		{Asset: "BTC", Balance: 0.23},
		{Asset: "ETH", Balance: 2.3},
		{Asset: "XMR", Balance: 43.145},
		{Asset: "ADA", Balance: 0.033},
	}, nil
}

func (exchangeKraken *ExchangeKraken) CreateOrder(exchangeConnection *Portfolio, asset string, amount float32) (CreatedOrder, error) {
	return CreatedOrder{
		OrderIdentifier: "123456",
	}, nil
}

func (exchangeKraken *ExchangeKraken) Holdings(exchangeConnection *Portfolio) ([]Holding, error) {
	return []Holding{}, nil
}
