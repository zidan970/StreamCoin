package handlers

import (
	"log"
	"zidan/gin-rest/entities"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	db, errOpen := gorm.Open(sqlite.Open("streamCoin.db"), &gorm.Config{})

	if errOpen != nil {
		// panic("failed to connect the database") // panic will make the next statemnt unreachable
		log.Fatal("failed to connect the database")
		return nil, errOpen
	}

	if errMigrate := MigrateSchema(db); errMigrate != nil {
		return nil, errMigrate
	}

	return db, nil
}

func Hello() {
	println("Hello World")
}

func MigrateSchema(db *gorm.DB) error {
	return db.AutoMigrate(&entities.Accounts{})
}
