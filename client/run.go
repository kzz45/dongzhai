package client

import (
	"dongzhai/models"
	"time"

	"gorm.io/gorm"
)

func dbCheck(db *gorm.DB) {
	var err error
	sql_db, err := db.DB()
	if err != nil {
		panic(err)
	}
	var count int
	for k := 0; k < 5; k++ {
		err = sql_db.Ping()
		if err != nil {
			count++
		}
		time.Sleep(time.Second)
		if count > 3 {
			panic(err)
		}
	}
	DBClient.AutoMigrate(&models.Cluster{})
}

func Run() {
	dbCheck(NewDBClient())
}
