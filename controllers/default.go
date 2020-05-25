package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	errJson := `{"code": 404 ,"message": "Resource / not found"}`
	c.Ctx.ResponseWriter.WriteHeader(404)
	c.Ctx.ResponseWriter.Write([]byte(errJson))
}
