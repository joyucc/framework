package orm

import (
	"sync"
	"time"

	"github.com/abaole/framework/logger"

	"github.com/jinzhu/gorm"
)

type DBHandler struct {
	DB *gorm.DB
}

func New(opts ...Option) *DBHandler {
	o := NewOptions(opts...)
	db := connect(o)
	db.DB().SetMaxIdleConns(o.Idle)
	db.DB().SetMaxOpenConns(o.Active)
	db.DB().SetConnMaxLifetime(time.Duration(o.IdleTimeout) / time.Second)
	db.SetLogger(ormLog{})
	return &DBHandler{db}
}

func (d *DBHandler) Close() {
	err := d.DB.Close()
	if err != nil {
		logger.Errorf("Disconnect from database failed: [%s]", err)
	}
	logger.Info("database closed")
}

var migratesOnce sync.Once

// migrate migrates database schemas ...
func (d *DBHandler) Migrate(models []interface{}) error {
	migratesOnce.Do(func() {
		err := d.DB.AutoMigrate(models...).Error
		if err != nil {
			logger.Panicf("auto migrate db table error: %v", err)
		}
	})

	return nil
}
