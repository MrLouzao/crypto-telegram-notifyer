package controllers

import (
	"crypto-telegram-notifyer/services"

	"github.com/astaxie/beego"
)

type HealthController struct {
	beego.Controller
}

func (this *HealthController) Get() {
	status := services.GetSystemHealth()

	this.Ctx.Output.IsOk()
	this.Ctx.Output.Header("Content-Type", "application/json;charset=UTF-8")
	this.Data["json"] = *status
	this.ServeJSON()
}
