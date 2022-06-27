package service

import "fmt"

type ExchangeMocked struct {
}

func (exchangeMocked *ExchangeMocked) SupportedAssets(portfolio *Portfolio) (map[string]*Asset, error) {
	return map[string]*Asset{
		"BTC":  FindAssetBySymbol("BTC"),
		"ETH":  FindAssetBySymbol("ETH"),
		"XMR":  FindAssetBySymbol("XMR"),
		"BNB":  FindAssetBySymbol("BNB"),
		"ADA":  FindAssetBySymbol("ADA"),
		"NANO": FindAssetBySymbol("NANO"),
	}, nil
}

// Gets the Exchange object for a given Exchange Connection, which is where the API call logic is.
func (portfolio *Portfolio) Exchange() (Exchange, error) {
	exchange, exists := exchanges[portfolio.ExchangeIdentifier]
	if !exists {
		return nil, fmt.Errorf("exchange %q is not implemented", portfolio.ExchangeIdentifier)
	}

	return exchange, nil
}

func (exchangeMocked *ExchangeMocked) CreateOrder(portfolio *Portfolio, asset string, amount float32) (CreatedOrder, error) {
	return CreatedOrder{
		OrderIdentifier: "123456",
	}, nil
}

func (exchangeMocked *ExchangeMocked) Holdings(portfolio *Portfolio) (map[string]Holding, error) {
	return map[string]Holding{
		"BTC": {Asset: FindAssetBySymbol("BTC"), Balance: 0.23},
		"ETH": {Asset: FindAssetBySymbol("ETH"), Balance: 2.3},
		"XMR": {Asset: FindAssetBySymbol("XMR"), Balance: 43.145},
		"BNB": {Asset: FindAssetBySymbol("BNB"), Balance: 0.033},
		"ADA": {Asset: FindAssetBySymbol("ADA"), Balance: 50.2},
	}, nil
}

func (exchangeMocked *ExchangeMocked) SupportsAsset(portfolio *Portfolio, asset Asset) bool {
	return true
}

func (exchangeMocked *ExchangeMocked) ValidateConnection(portfolio *Portfolio) ValidateExchangeConnectionResult {
	return ValidateExchangeConnectionResult{
		Success: true,
		Issue:   "",
	}
}

func (exchangeMocked *ExchangeMocked) HoldingSummary(portfolio *Portfolio) (HoldingSummary, error) {
	return HoldingSummary{
		TotalBalanceValuation: 1234.56,
	}, nil
}
