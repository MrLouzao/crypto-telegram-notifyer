package main

import (
	"crypto-telegram-notifyer/conf"
	_ "crypto-telegram-notifyer/routers"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	go conf.StartDatabase()
	beego.Run()
}
