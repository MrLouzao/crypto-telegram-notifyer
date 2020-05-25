package services

import (
	"crypto-telegram-notifyer/models"

	_ "crypto-telegram-notifyer/models"
	"database/sql"
	"fmt"

	"github.com/astaxie/beego/orm"
)

type alarmResults struct {
	items []models.Alarm
}

// Obtain all alarms from DB
func query_all_alarms(db *sql.DB) *alarmResults {
	o := orm.NewOrm()
	qs := o.QueryTable("alarm")
	var result []models.Alarm
	qs.All(&result)

	alarmEntities := alarmResults{items: result}
	return &alarmEntities
}

// Cast from entities to dto
func cast_to_dto(alarmEntities *alarmResults) *[]models.AlarmDto {
	entities := alarmEntities.items
	dtos := make([]models.AlarmDto, len(entities))

	for i := range entities {
		dtos[i] = models.AlarmDto(entities[i])
	}

	return &dtos
}

// Return a list of stored alarms on DB
func GetAllAlarms() ([]models.AlarmDto, error) {
	db, err := orm.GetDB("default")
	if err != nil {
		fmt.Println(err)
		return nil, err
	} else {
		res := query_all_alarms(db)
		resDto := cast_to_dto(res)
		return *resDto, nil
	}
}

// Insert a new row on of alarms
func insert_new_alarm(entity models.Alarm) error {
	o := orm.NewOrm()
	o.Using("default")

	fmt.Println("EL ORM ES: ", o)
	// Insert requires the address of the object, not the value!!!
	res, err := o.Insert(&entity)
	if err != nil {
		return err
	} else {
		fmt.Println(res)
		return nil
	}
}

// Save a new Alarm
func SaveAlarm(dto models.AlarmDto) error {
	_, err := orm.GetDB("default")
	if err != nil {
		fmt.Println(err)
		return err
	} else {
		entity := models.Alarm(dto)
		err := insert_new_alarm(entity)
		return err
	}
}
