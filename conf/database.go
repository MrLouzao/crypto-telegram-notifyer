package conf

import (
	"crypto-telegram-notifyer/models"
	"fmt"

	"os"

	"github.com/astaxie/beego/orm"
)

func recreate_db(dbAlias string) {
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

// Initialize DB Setup
func StartDatabase() {
	// Startup DB
	dbAlias := "default"
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase(dbAlias, "mysql", "root:example@/crypto_notifier?charset=utf8")
	orm.SetMaxIdleConns("default", 30)
	orm.SetMaxOpenConns("default", 30)
	orm.RunCommand()

	recreateDbEnv := os.Getenv("RECREATE_DB")
	if recreateDbEnv == "true" {
		defer addMockedData()
		recreate_db(dbAlias)
	}

}

func addMockedData() {
	o := orm.NewOrm()
	o.Using("default")

	alarm := new(models.Alarm)
	alarm.Id = 1
	alarm.Name = "bitisi"
	alarm.Type = "UP"
	alarm.Against = "btc"
	alarm.Price = 1232

	fmt.Println(o.Insert(alarm))
}
