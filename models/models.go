package models

import (
	"github.com/astaxie/beego/orm"
)

func init() {
	// Need to register model in init
	orm.RegisterModel(new(Alarm))
}
