package database

import (
	"github.com/abaole/framework/database/orm"
	"github.com/jinzhu/gorm"
)

//默认
var l IDatabase

type IDatabase interface {
	Close()
}

//设置
func SetDatabase(ll IDatabase) {
	l = ll
}

func Close() {
	l.Close()
}

func GetOrm() *gorm.DB {
	return orm.GetDB()
}
