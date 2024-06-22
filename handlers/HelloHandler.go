package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	DB   *gorm.DB
	side float32
}

func InitHandler(db *gorm.DB) *Handler {
	var s float32 = 5
	return &Handler{DB: db, side: s}
}

func (h *Handler) TestHello(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello World")
}

func (h *Handler) TestCount(c *gin.Context) { // we can't make this handler have a return other than void or error
	count := h.side * h.side
	c.JSON(http.StatusOK, gin.H{
		"count": count,
	})
}
