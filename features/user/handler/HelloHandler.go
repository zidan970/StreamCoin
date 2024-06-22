package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// in order to use db in every method, i use the handler (that contains db) in every method
// in order to use service in every method, i use the handler (that contains service) in every method
type UserHandler struct {
	DB   *gorm.DB
	side float32
}

func InitHandler(db *gorm.DB) *UserHandler {
	var s float32 = 5
	return &UserHandler{DB: db, side: s}
}

// type UserHandler struct {
// 	userService user.service
// 	side        float32
// }

// func InitHandler(userService user.service) *UserHandler {
// 	return &UserHandler{userService: userService}
// }

func (h *UserHandler) TestHello(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello World")
}

func (h *UserHandler) TestCount(c *gin.Context) { // we can't make this handler have a return other than void or error
	count := h.side * h.side
	c.JSON(http.StatusOK, gin.H{
		"count": count,
	})
}
