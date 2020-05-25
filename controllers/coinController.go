package controllers

import (
	"errors"
	"regexp"
	"strings"

	"crypto-telegram-notifyer/coingecko"

	"github.com/astaxie/beego"
)

type CoinController struct {
	beego.Controller
}

// Definition of a response with data
type CoinResponse struct {
	Name     string `json:"symbol"`
	UsdPrice string `json:"usd_price"`
	BtcPrice string `json:"btc_price"`
}

type ErrorMsg struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
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

// For each coin, retrieve a CoinResponse object
func _obtain_coin_list_prices(coinList []string) (*[]CoinResponse, error) {
	coinPricesList, err := coingecko.GetCoinsPrices(coinList)

	if err != nil {
		return nil, errors.New("Error while calling to Coingecko's API")
	} else {

		var coinPricesListToReturn = make([]CoinResponse, len(*coinPricesList))

		for i, coinPriceResponse := range *coinPricesList {
			// Cast between interfaces
			coinResponseCasted := CoinResponse(coinPriceResponse)
			coinPricesListToReturn[i] = coinResponseCasted
		}

		return &coinPricesListToReturn, nil
	}
}

// Response for GET:/coins?symbols=btc,lsk
func (this *CoinController) Get() {
	coinSymbols := this.GetString("symbols")
	coinSymbolList, err := _extract_coin_symbols(coinSymbols)

	// If list of coins is not well formed send a 403 error
	if err != nil {
		_send_403_error_response(this.Ctx.ResponseWriter, err.Error())
	} else {

		coinPricesList, err2 := _obtain_coin_list_prices(coinSymbolList)
		if err2 != nil {
			_send_500_error_response(this.Ctx.ResponseWriter, err2.Error())
		} else {
			this.Data["json"] = *coinPricesList
			this.ServeJSON()
		}

	}
}
