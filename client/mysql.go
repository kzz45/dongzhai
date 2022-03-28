package client

import (
	"dongzhai/config"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DBClient *gorm.DB

func NewDBClient() *gorm.DB {
	if DBClient == nil {
		db, err := gorm.Open(mysql.New(mysql.Config{
			DSN: fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
				config.GlobalConfig.Database.Username,
				config.GlobalConfig.Database.Password,
				config.GlobalConfig.Database.Host,
				config.GlobalConfig.Database.Port,
				config.GlobalConfig.Database.Name,
			),
		}), &gorm.Config{
			Logger:                                   logger.Default.LogMode(logger.Info),
			DisableForeignKeyConstraintWhenMigrating: true,
		})
		if err != nil {
			panic(err)
		}
		DBClient = db
		return DBClient
	}
	return DBClient
}