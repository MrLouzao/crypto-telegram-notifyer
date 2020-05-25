package conf

import (
	"crypto-telegram-notifyer/models"
	"fmt"

	"github.com/astaxie/beego/orm"
)

// Initialize DB Setup
func StartDatabase() {
	// Startup DB
	dbAlias := "default"
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase(dbAlias, "mysql", "root:example@/crypto_notifier?charset=utf8")

	// Create DB with tables
	//defer addMockedData()
	orm.RunCommand()

	// Drop table and recreate
	force := true
	// Print log
	verbose := true
	// Sync and check if fails
	err := orm.RunSyncdb(dbAlias, force, verbose)
	if err != nil {
		fmt.Println(err)
		panic(1)
	}
}

func addMockedData() {
	o := orm.NewOrm()
	o.Using("default")

	alarm := new(models.Alarm)
	alarm.Id = 1
	alarm.Name = "bitisi"
	alarm.Against = "btc"
	alarm.Price = 1232

	fmt.Println(o.Insert(alarm))
}
