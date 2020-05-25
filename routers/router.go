package routers

import (
	"crypto-telegram-notifyer/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/health", &controllers.HealthController{})
	beego.Router("/coins", &controllers.CoinController{})
	beego.Router("/alarms", &controllers.AlarmController{})
}
