package services

import (
	"crypto-telegram-notifyer/coingecko"
	"crypto-telegram-notifyer/mappers"
	"crypto-telegram-notifyer/models"
	"crypto-telegram-notifyer/repositories"
	"errors"
	"fmt"

	_ "crypto-telegram-notifyer/models"
	_ "crypto-telegram-notifyer/repositories"
)

type alarmResults struct {
	items []models.Alarm
}

// Return a list of stored alarms on DB
func GetAllAlarms() ([]models.AlarmDto, error) {
	allAlarms, err := repositories.QueryAllAlarms()
	if err != nil {
		return nil, err
	} else {
		resDto := mappers.CastToAlarmDtoList(allAlarms)
		return *resDto, nil
	}
}

// Check if a coin pair is available based on Coingeckos API response
func check_coin_pair_exists_on_api(dto models.AlarmDto) bool {
	_, err := coingecko.GetCoinPriceAgainst(dto.Name, dto.Against)
	if err != nil {
		return false
	} else {
		return true
	}
}

// Save a new Alarm
func SaveAlarm(dto models.AlarmDto) error {
	// First check if the alarm is available
	if dto.Against != coingecko.USD_SYMBOL && dto.Against != coingecko.BTC_SYMBOL {
		return errors.New("Only BTC and USD against pairs supported")
	}

	availableOnApi := check_coin_pair_exists_on_api(dto)
	if !availableOnApi {
		return errors.New("Given pair is not supported by Coingeckos API")
	}

	entity := models.Alarm(dto)
	err := repositories.SaveNewAlarm(entity)
	return err
}

// Perform all alarms by checking prices on Coingeckos DB
func CheckAlarmsThroughCoingeckoApi() {
	allAlarms, err := repositories.QueryAllAlarms()
	if err != nil {
		fmt.Println("Error while obtaining all configured alarms")
	} else {
		// TODO implement this operation
		// 1. Map from records to pairs of coins
		// 2. Check against API
		// 3. Map results to readable format
		// 4. Check each alarm and notify
		fmt.Println("Alarms configured", allAlarms)
	}
}
