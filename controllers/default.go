package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "luismiguel.louzao.gonzalez@gmail.com"
	c.TplName = "index.tpl"
}
