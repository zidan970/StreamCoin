package main

import (
	"log"
	"zidan/gin-rest/apps"
	"zidan/gin-rest/features/user/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	// init db
	db, errInit := apps.InitDB()
	if errInit != nil {
		log.Fatal("db initialization is failed: ", errInit)
	}

	// init handler
	hr := handler.InitHandler(db)

	// init gin
	r := gin.New()

	r.GET("/hello", hr.TestHello)
	// keep in mind that in the Gin framework, handlers for routes must have the appropriate format
	// that fits gin.HandlerFunc, which has type SIGNATURE func(*gin.Context).
	// This means that the handler must accept the *gin.Context parameter and return VOID or ERROR.
	r.GET("/count", hr.TestCount)
	r.POST("/users", hr.Register)

	r.Run()
}

// old school: var db, db used in parameter
// modern: db used in parameter init handler, return a struct that have db as a field, use that struct for every function (now it's called method)

// type Handler struct {
// 	db *gorm.DB
// }

// func newHandler(db *gorm.DB) *Handler {
// 	return &Handler{db}
// }
