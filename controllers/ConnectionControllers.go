package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func InitDB() (*sql.DB, error) {
	// open connection
	db, errOpen := sql.Open("mysql", os.Getenv("AppConnection"))
	// <username>:<password>@tcp(<hostname>:<port-db>)/<db-name>
	// db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/db_elibrary?parseTime=true")
	if errOpen != nil {
		log.Fatal("open connection is failed: ", errOpen)
		return nil, errOpen
	}

	// test connection
	errPing := db.Ping()
	if errPing != nil {
		log.Fatal("no connection available: ", errPing)
		return nil, errPing
	} else {
		fmt.Println("you are succesfully connected to database")
		return db, nil
	}
}
