package db

import (
	"dongzhai/config"
	"dongzhai/models"
	"dongzhai/models/k8s"
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
	GlobalGorm.AutoMigrate(&models.Role{})
	GlobalGorm.AutoMigrate(&models.UserGroup{})

	GlobalGorm.AutoMigrate(&models.Cloud{})
	GlobalGorm.AutoMigrate(&models.Product{})

	// GlobalGorm.AutoMigrate(&domain.Domain{})
	// GlobalGorm.AutoMigrate(&domain.DomainCert{})
	// GlobalGorm.AutoMigrate(&domain.DomainRecord{})

	// GlobalGorm.AutoMigrate(&monitor.Task{})
	// GlobalGorm.AutoMigrate(&monitor.Label{})
	// GlobalGorm.AutoMigrate(&monitor.Server{})
	// GlobalGorm.AutoMigrate(&monitor.Receiver{})
	// GlobalGorm.AutoMigrate(&monitor.Instance{})
	// GlobalGorm.AutoMigrate(&monitor.AlertRule{})
	// GlobalGorm.AutoMigrate(&monitor.LabelValue{})
	// GlobalGorm.AutoMigrate(&monitor.AlertRoute{})

	GlobalGorm.AutoMigrate(&k8s.Pod{})
	GlobalGorm.AutoMigrate(&k8s.Job{})
	GlobalGorm.AutoMigrate(&k8s.Node{})
	GlobalGorm.AutoMigrate(&k8s.Secret{})
	GlobalGorm.AutoMigrate(&k8s.Project{})
	GlobalGorm.AutoMigrate(&k8s.Cluster{})
	GlobalGorm.AutoMigrate(&k8s.Service{})
	GlobalGorm.AutoMigrate(&k8s.Ingress{})
	GlobalGorm.AutoMigrate(&k8s.Registry{})
	GlobalGorm.AutoMigrate(&k8s.Container{})
	GlobalGorm.AutoMigrate(&k8s.ConfigMap{})
	GlobalGorm.AutoMigrate(&k8s.Deployment{})
	GlobalGorm.AutoMigrate(&k8s.ServicePort{})
	GlobalGorm.AutoMigrate(&k8s.IngressRule{})
	GlobalGorm.AutoMigrate(&k8s.ContainerPort{})
}
