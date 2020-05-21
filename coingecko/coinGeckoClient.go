package coingecko

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"time"

	coingecko "github.com/superoo7/go-gecko/v3"
)

// Constants
const BTC_SYMBOL string = "btc"
const USD_SYMBOL string = "usd"

type CoinPriceResponse struct {
	Name     string
	UsdPrice string
	BtcPrice string
}

func ApiIsUp() bool {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}
	coinGeckoApiClient := coingecko.NewClient(httpClient)

	_, err := coinGeckoApiClient.Ping()

	if err != nil {
		fmt.Println("Error while ping Coingecko's API", err.Error())
		return false
	} else {
		return true
	}
}

// From Map of counterparty values to CoinPriceResponse
func _mapResponseToCoinPriceResponse(coinName string, coinPrices *map[string]float32) CoinPriceResponse {
	btcPrice := (*coinPrices)[BTC_SYMBOL]
	usdPrice := (*coinPrices)[USD_SYMBOL]
	coinPrice := CoinPriceResponse{
		Name:     strings.ToUpper(coinName),
		UsdPrice: fmt.Sprintf("%f", usdPrice),
		BtcPrice: fmt.Sprintf("%f", btcPrice),
	}
	return coinPrice
}

// Cast response from CoinGecko API to a list of CoinPriceResponse structs
func mapCoinResultToCoinPriceResponse(res *map[string]map[string]float32) *[]CoinPriceResponse {
	coinSymbolList := reflect.ValueOf(*res).MapKeys()
	coinPricesToReturn := make([]CoinPriceResponse, len(coinSymbolList))

	for i, coinSymbolKey := range coinSymbolList {
		coinSymbol := coinSymbolKey.String()
		coinPriceMap := (*res)[coinSymbol]
		castedCoinPrice := _mapResponseToCoinPriceResponse(coinSymbol, &coinPriceMap)
		coinPricesToReturn[i] = castedCoinPrice
	}

	return &coinPricesToReturn
}

// Obtain a Ticker price from a given symbol
func GetCoinsPrices(inputSymbols []string) (*[]CoinPriceResponse, error) {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}
	coinGeckoApiClient := coingecko.NewClient(httpClient)

	// Hardcode counterparty symbols
	againstSymbols := []string{BTC_SYMBOL, USD_SYMBOL}
	res, err := coinGeckoApiClient.SimplePrice(inputSymbols, againstSymbols)

	if err != nil {
		return nil, err
	} else {
		var castedCoinPriceResponseList = mapCoinResultToCoinPriceResponse(res)
		return castedCoinPriceResponseList, nil
	}
}
