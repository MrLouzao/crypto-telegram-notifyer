package controllers

import (
	"errors"
	"net/http"
	"regexp"
	"strings"

	"github.com/astaxie/beego"
)

type CoinController struct {
	beego.Controller
}

// Definition of a response with data
type CoinResponse struct {
	Symbol   string `json:"symbol"`
	UsdPrice string `json:"usd_price"`
	BtcPrice string `json:"btc_price"`
}

type ErrorMsg struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Method to call to CoinGecko API and retrieve data
func _obtain_coin_price(symbol *string) CoinResponse {
	// TODO call to the correspondent API
	mockBtcPrice := CoinResponse{
		Symbol:   strings.ToUpper(*symbol),
		UsdPrice: "100k",
		BtcPrice: "1.00000000000",
	}
	return mockBtcPrice
}

// Extract all coins from array splitting by comma
func _extract_coin_symbols(allCoinsCommaSeparated string) ([]string, error) {
	coinListStrIsNotEmpty := len(allCoinsCommaSeparated) > 0

	if coinListStrIsNotEmpty {

		match, _ := regexp.MatchString("([a-z]+)(,*)", allCoinsCommaSeparated)

		if !match {
			return nil, errors.New("Coin list must be separated by commas. Not other symbols or numbers allowed")
		} else {
			// Only return all coin list when is ready
			var coinList []string = strings.Split(allCoinsCommaSeparated, ",")
			return coinList, nil
		}

	} else {
		return nil, errors.New("Coin symbol list must not be empty")
	}
}

// Generate a 403 error from Response
func _send_403_error_response(writer http.ResponseWriter, errMessage string) {
	errJson := `{"code": 403 ,"message": "` + errMessage + `"}`
	writer.WriteHeader(403)
	writer.Write([]byte(errJson))
}

// For each coin, retrieve a CoinResponse object
func _obtain_coin_list_prices(coinList []string) *[]CoinResponse {
	var coinPricesList = make([]CoinResponse, len(coinList))

	for i, coinSymbol := range coinList {
		coinPricesList[i] = _obtain_coin_price(&coinSymbol)
	}

	return &coinPricesList
}

// Response for GET:/coins?symbols=btc,lsk
func (this *CoinController) Get() {
	coinSymbols := this.GetString("symbols")
	coinSymbolList, err := _extract_coin_symbols(coinSymbols)

	// If list of coins is not well formed send a 403 error
	if err != nil {
		_send_403_error_response(this.Ctx.ResponseWriter, err.Error())
	} else {
		coinPricesList := _obtain_coin_list_prices(coinSymbolList)
		this.Data["json"] = *coinPricesList
		this.ServeJSON()
	}
}
