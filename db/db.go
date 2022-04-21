package db

import (
	"dongzhai/config"
	"dongzhai/models"
	"dongzhai/models/domain"
	"dongzhai/models/monitor"
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var GlobalGorm *gorm.DB

func init() {
	var err error
	log_level := logger.Warn
	if config.GlobalConfig.Server.Debug {
		log_level = logger.Info
	}
	GlobalGorm, err = gorm.Open(mysql.New(mysql.Config{
		DSN: fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			config.GlobalConfig.Database.Username,
			config.GlobalConfig.Database.Password,
			config.GlobalConfig.Database.Host,
			config.GlobalConfig.Database.Port,
			config.GlobalConfig.Database.Name,
		),
	}), &gorm.Config{
		Logger:                                   logger.Default.LogMode(log_level),
		DisableForeignKeyConstraintWhenMigrating: false,
		PrepareStmt:                              true,
	})

	if err != nil {
		logrus.Fatalln(err)
	}
	logrus.Infoln("connect mysql success")

	GlobalGorm.AutoMigrate(&models.User{})
	GlobalGorm.AutoMigrate(&models.UserGroup{})

	GlobalGorm.AutoMigrate(&models.Cloud{})
	GlobalGorm.AutoMigrate(&models.Product{})

	GlobalGorm.AutoMigrate(&domain.Domain{})
	GlobalGorm.AutoMigrate(&domain.DomainCert{})
	GlobalGorm.AutoMigrate(&domain.DomainRecord{})

	GlobalGorm.AutoMigrate(&monitor.Task{})
	GlobalGorm.AutoMigrate(&monitor.Label{})
	GlobalGorm.AutoMigrate(&monitor.Server{})
	GlobalGorm.AutoMigrate(&monitor.Receiver{})
	GlobalGorm.AutoMigrate(&monitor.Instance{})
	GlobalGorm.AutoMigrate(&monitor.AlertRule{})
	GlobalGorm.AutoMigrate(&monitor.LabelValue{})
	GlobalGorm.AutoMigrate(&monitor.AlertRoute{})
}
