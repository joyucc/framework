package orm

import (
	"strings"
	"time"

	"github.com/abaole/framework/logger"

	"github.com/jinzhu/gorm"
)

func connect(o *Options) *gorm.DB {
	db, err := gorm.Open(o.Dialect, o.DSN)
	if err != nil {
		panic("链接数据库失败")
	}

	// 数据库心跳测试
	if err = pingDatabase(db); err != nil {
		panic("数据库ping不通")
	}

	return db
}

// helper function to ping the database with backoff to ensure
// a connection can be established before we proceed with the
func pingDatabase(g *gorm.DB) (err error) {
	for i := 0; i < 30; i++ {
		err = g.DB().Ping()
		if err == nil {
			return
		}
		time.Sleep(time.Second)
	}
	return
}

type ormLog struct{}

func (l ormLog) Print(v ...interface{}) {
	logger.Info(strings.Repeat("%v ", len(v)), v...)
}
