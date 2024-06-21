package main

import (
	"fmt"
	"log"
	"zidan/AccountServiceAppProject/controllers"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// make connection
	db, errInitDB := controllers.InitDB()
	if errInitDB != nil {
		log.Fatal("init db is failed: ", errInitDB)
	}

	defer db.Close()

	fmt.Println()

	controllers.Menu(db)
}
