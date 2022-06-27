package service

import (
	"fmt"
	"log"
	"strconv"

	krakenapi "github.com/beldur/kraken-go-api-client"
)

type ExchangeKraken struct {
}

func (exchangeKraken *ExchangeKraken) CreateOrder(exchangeConnection *Portfolio, asset string, amount float32) (CreatedOrder, error) {
	return CreatedOrder{
		OrderIdentifier: "123456",
	}, nil
}

func (exchangeKraken *ExchangeKraken) Holdings(exchangeConnection *Portfolio) (map[string]Holding, error) {
	return map[string]Holding{}, nil
}

func (exchangeKraken *ExchangeKraken) SupportsAsset(exchangeConnection *Portfolio, asset Asset) bool {
	return true
}

func (exchangeKraken *ExchangeKraken) ValidateConnection(portfolio *Portfolio) ValidateExchangeConnectionResult {
	api := krakenapi.New(portfolio.ApiKey, portfolio.ApiSecret)
	_, err := api.Query("TradeBalance", map[string]string{
		"asset": "USD",
	})

	if err != nil {
		return ValidateExchangeConnectionResult{
			Success: false,
			Issue:   fmt.Sprintf("could not query Kraken: %s", err),
		}
	}

	return ValidateExchangeConnectionResult{
		Success: true,
		Issue:   "",
	}
}

func (exchangeKraken *ExchangeKraken) SupportedAssets(portfolio *Portfolio) (map[string]*Asset, error) {
	supportedAssets := make(map[string]*Asset, 0)

	api := krakenapi.New(portfolio.ApiKey, portfolio.ApiSecret)
	assetsResponseRaw, err := api.Query("Assets", map[string]string{})
	if err != nil {
		return supportedAssets, err
	}

	assetsResponse := assetsResponseRaw.(map[string]interface{})

	for symbol, assetRaw := range assetsResponse {
		var altSymbol string
		exchangeAsset := assetRaw.(map[string]interface{})
		foundAsset := FindAssetBySymbol(symbol)

		// If asset is not found, check by `altname`.
		if foundAsset == nil {
			altSymbolRaw, hasAltSymbol := exchangeAsset["altname"]

			// An alt name is provided, try finding the asset by alt symbol.
			if hasAltSymbol {
				altSymbol = altSymbolRaw.(string)
				foundAsset = FindAssetBySymbol(altSymbol)
			}
		}

		// Only add to the list of supported assets if we were able to find an asset.
		if foundAsset != nil {
			supportedAssets[foundAsset.Symbol] = foundAsset
		} else {
			log.Printf("No asset found for symbol: %s", symbol)

			if len(altSymbol) > 0 {
				log.Printf("No asset found for alt symbol: %s", symbol)
			}
		}
	}

	return supportedAssets, nil
}

func (exchangeKraken *ExchangeKraken) HoldingSummary(portfolio *Portfolio) (HoldingSummary, error) {
	api := krakenapi.New(portfolio.ApiKey, portfolio.ApiSecret)
	tradeBalancesResponseRaw, err := api.Query("TradeBalance", map[string]string{
		"asset": "USD",
	})
	if err != nil {
		return HoldingSummary{}, err
	}

	tradeBalanceResponse := tradeBalancesResponseRaw.(map[string]interface{})

	totalBalance, err := strconv.ParseFloat(tradeBalanceResponse["eb"].(string), 64)

	return HoldingSummary{
		TotalBalanceValuation: totalBalance,
	}, nil
}
