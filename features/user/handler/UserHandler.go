package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *UserHandler) Register(c *gin.Context) {
	newRegister := UserDTO{}

	// binding data
	if err := c.ShouldBindJSON(&newRegister); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erorr bind": err.Error(),
		})
		return
	}

	newRegisModel := UserDTOtoModel(newRegister)

	// we save/create the binded request to db
	// user.UserId = uuid.New().String()
	if result := h.DB.Create(&newRegisModel); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error create": result.Error.Error(),
		})
	}

}
