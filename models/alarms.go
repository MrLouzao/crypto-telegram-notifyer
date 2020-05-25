package models

import "errors"

// Define a configurable alarm
type Alarm struct {
	Id      int
	Name    string
	Type    string
	Against string
	Price   int
}

// DTO for alarms
type AlarmDto struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	Against string `json:"against"`
	Price   int    `json:"price"`
}

// Check json format
func CheckPostAlarmDtoFormat(alarm *AlarmDto) error {
	if alarm.Id > 0 {
		return errors.New("Alarm Id must not be set when creating new")
	}
	if len(alarm.Name) == 0 {
		return errors.New("Coin Name must be set when creating new")
	}
	if len(alarm.Against) == 0 {
		return errors.New("You must set the coin against to compare (btc or usd)")
	}
	if len(alarm.Type) == 0 {
		return errors.New("Must configure the type of alarm (below/above price)")
	}
	if alarm.Price < 0 {
		return errors.New("Target price must be set when creating new")
	}

	return nil
}
