package repositories

import (
	"crypto-telegram-notifyer/models"
	"database/sql"
	"fmt"

	"github.com/astaxie/beego/orm"
)

type AlarmResults struct {
	Items []models.Alarm
}

// Obtain all alarms from DB
func query_all_alarms(db *sql.DB) *AlarmResults {
	o := orm.NewOrm()
	qs := o.QueryTable("alarm")
	var result []models.Alarm
	qs.All(&result)

	alarmEntities := AlarmResults{Items: result}
	return &alarmEntities
}

// Query all stored alarms
func QueryAllAlarms() (*AlarmResults, error) {
	db, err := orm.GetDB("default")
	if err != nil {
		fmt.Println(err)
		return nil, err
	} else {
		res := query_all_alarms(db)
		return res, nil
	}
}

// Insert a new row on of alarms
func SaveNewAlarm(entity models.Alarm) error {
	o := orm.NewOrm()
	o.Using("default")

	// Insert requires the address of the object, not the value!!!
	res, err := o.Insert(&entity)
	if err != nil {
		return err
	} else {
		fmt.Println(res)
		return nil
	}
}
