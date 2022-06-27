package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/jaksonkallio/coinbake/config"
	"github.com/jaksonkallio/coinbake/database"
)

type ListingsResponse struct {
	Listings []Listing `json:"data"`
}

type Listing struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
	Quotes Quotes `json:"quote"`
}

type Quotes struct {
	USD Quote `json:"USD"`
}

type Quote struct {
	Price     float64 `json:"price"`
	Volume    float64 `json:"volume_24h"`
	MarketCap float64 `json:"market_cap"`
}

const (
	CaptureTopAssetCount int = 200
)

func MarketDataRefresher(ticks *time.Ticker, stop chan bool) {
	// Initial market data refresh.
	refreshMarketData()

	for {
		select {
		case <-stop:
			return
		case <-ticks.C:
			refreshMarketData()
		}
	}
}

// Refreshes all market data.
func refreshMarketData() {
	log.Println("Refreshing market data")

	responseBytes, err := marketDataApi("cryptocurrency/listings/latest", url.Values{
		"start":   []string{"1"},
		"limit":   []string{fmt.Sprintf("%d", CaptureTopAssetCount)},
		"convert": []string{"USD"},
	})
	if err != nil {
		log.Printf("Could not reach market data API: %s", err)
	}

	listingsResponse := ListingsResponse{}
	json.Unmarshal(responseBytes, &listingsResponse)

	for _, listing := range listingsResponse.Listings {
		asset := FindAssetBySymbol(listing.Symbol)
		if asset == nil {
			asset = CreateAsset(listing.Symbol)
		}

		asset.MarketCap = uint64(listing.Quotes.USD.MarketCap)
		asset.Volume = uint64(listing.Quotes.USD.Volume)
		asset.ApproxPrice = listing.Quotes.USD.Price
		asset.LastRefreshed = time.Now()
		asset.Name = listing.Name
		database.Handle().Save(asset)
	}
}

func marketDataApi(endpoint string, query url.Values) ([]byte, error) {
	if len(config.CurrentConfig.MarketData.ApiKey) == 0 {
		return nil, fmt.Errorf("did not configure a market data API key")
	}

	if len(config.CurrentConfig.MarketData.BaseUrl) == 0 {
		return nil, fmt.Errorf("did not configure a market data API base URL")
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", config.CurrentConfig.MarketData.BaseUrl, endpoint), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", config.CurrentConfig.MarketData.ApiKey)
	req.URL.RawQuery = query.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}
