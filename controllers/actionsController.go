package controllers

import (
	"crypto-telegram-notifyer/services"

	"github.com/astaxie/beego"
)

type ActionsController struct {
	beego.Controller
}

// Check all programmed alarms POST:/actions
func (this *ActionsController) Post() {
	writer := this.Ctx.ResponseWriter
	go services.CheckAlarmsThroughCoingeckoApi()
	writer.WriteHeader(201)
}
