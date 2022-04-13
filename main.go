package main

import (
	"dongzhai/config"
	"dongzhai/db"
	"fmt"
)

func main() {
	db, _ := db.GlobalGorm.DB()
	err := db.Ping()
	fmt.Println(err)

	fmt.Println(config.GlobalConfig.Server.Name)
}
