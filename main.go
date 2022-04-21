package main

import (
	"dongzhai/db"
	"dongzhai/router"
	"fmt"
)

func main() {
	db, _ := db.GlobalGorm.DB()
	err := db.Ping()
	fmt.Println(err)

	router.InitRouter()
}
