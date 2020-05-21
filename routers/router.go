package routers

import (
	"crypto-telegram-notifyer/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/coins", &controllers.CoinController{})
}
