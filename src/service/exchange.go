package service

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

type SupportedAsset struct {
	Asset Asset
}

// An interface representing a generic exchange.
type Exchange interface {
	CreateOrder(*Portfolio, string, float32) (CreatedOrder, error)
	Holdings(*Portfolio) (map[string]Holding, error)
	SupportedAssets(*Portfolio) (map[string]*Asset, error)
	ValidateConnection(*Portfolio) ValidateExchangeConnectionResult
	HoldingSummary(*Portfolio) (HoldingSummary, error)
}

type ValidateExchangeConnectionResult struct {
	// Whether we're able to connect to the exchange at all.
	Success bool

	// Any error message that we may want to provide to the user about the connection.
	Issue string
}

type MockSupportedAssets struct {
}

type CreatedOrder struct {
	OrderIdentifier string
}

type Holding struct {
	Asset   *Asset
	Balance float64
}

type HoldingSummary struct {
	// How much the user holds in total.
	TotalBalanceValuation float64
}
