package controllers

import (
	"crypto-telegram-notifyer/models"
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
)

type AlarmController struct {
	beego.Controller
}

// Create an Alarm POST:/alarms
func (this *AlarmController) Post() {
	// Get a decoders
	req := *this.Ctx.Request
	decoder := json.NewDecoder(req.Body)

	// Cast from payload to alarm
	var alarm models.AlarmDto
	err := decoder.Decode(&alarm)
	if err != nil {
		fmt.Println(err.Error())
		_send_403_error_response(this.Ctx.ResponseWriter, "Alarm message format invalid")
	}

	// Check format
	err = models.CheckPostAlarmDtoFormat(&alarm)
	if err != nil {
		_send_403_error_response(this.Ctx.ResponseWriter, err.Error())
	} else {
		fmt.Println("alarm", alarm)
		writer := this.Ctx.ResponseWriter
		errMessage := " Method not ready yet!"
		errJson := `{"code": 403 ,"message": "` + errMessage + `"}`
		writer.WriteHeader(403)
		writer.Write([]byte(errJson))
	}
}

// Return all stored alarms GET:/alarms
func (this *AlarmController) Get() {
	var emptyArray = []string{}

	this.Data["json"] = emptyArray
	this.ServeJSON()
}
