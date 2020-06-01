package controllers

import (
	"crypto-telegram-notifyer/models"
	"crypto-telegram-notifyer/services"
	_ "crypto-telegram-notifyer/services"
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
)

type AlarmController struct {
	beego.Controller
}

// Create an Alarm POST:/alarms
func (this *AlarmController) Post() {
	writer := this.Ctx.ResponseWriter
	// Get a decoders
	req := *this.Ctx.Request
	decoder := json.NewDecoder(req.Body)

	// Cast from payload to alarm
	var alarm models.AlarmDto
	err := decoder.Decode(&alarm)
	if err != nil {
		fmt.Println(err.Error())
		_send_403_error_response(writer, "Alarm message format invalid")
	}

	// Check format
	err = models.CheckPostAlarmDtoFormat(&alarm)
	if err != nil {
		_send_403_error_response(writer, err.Error())
	} else {
		err := services.SaveAlarm(alarm)
		if err != nil {
			_send_500_error_response(writer, err.Error())
		} else {
			writer.WriteHeader(201)
		}
	}
}

// Return all stored alarms GET:/alarms
func (this *AlarmController) Get() {
	res, err := services.GetAllAlarms()
	if err != nil {
		writer := this.Ctx.ResponseWriter
		errMessage := "Cannot connect to DB!"
		errJson := `{"code": 500 ,"message": "` + errMessage + `"}`
		writer.WriteHeader(500)
		writer.Write([]byte(errJson))
	} else {
		this.Data["json"] = res
		this.ServeJSON()
	}
}

// Delete a stored alarm  DELETE:/alarms/{alarmId}
func (this *AlarmController) Delete() {
	fmt.Println("TODO: implement this endpoint")
	this.Ctx.ResponseWriter.WriteHeader(200)
}
