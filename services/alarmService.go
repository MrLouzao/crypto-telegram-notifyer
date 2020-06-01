package services

import (
	"crypto-telegram-notifyer/mappers"
	"crypto-telegram-notifyer/models"
	"crypto-telegram-notifyer/repositories"

	_ "crypto-telegram-notifyer/models"
	_ "crypto-telegram-notifyer/repositories"
	"fmt"

	"github.com/astaxie/beego/orm"
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

// Save a new Alarm
func SaveAlarm(dto models.AlarmDto) error {
	// TODO Check if alarm exists though API

	_, err := orm.GetDB("default")
	if err != nil {
		fmt.Println(err)
		return err
	} else {
		entity := models.Alarm(dto)
		err := repositories.SaveNewAlarm(entity)
		return err
	}
}
