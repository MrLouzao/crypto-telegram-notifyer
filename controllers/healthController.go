package controllers

import "github.com/astaxie/beego"

type HealthController struct {
	beego.Controller
}

func (this *HealthController) Get() {
	this.Ctx.ResponseWriter.WriteHeader(200)
	okMsg := "Server up and running!"
	this.Ctx.ResponseWriter.Write([]byte(okMsg))
}
