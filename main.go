package main

import (
	"crypto-telegram-notifyer/conf"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	conf.StartDatabase()
	beego.Run()
}
