package handler

import (
	"zidan/gin-rest/apps"

	"github.com/google/uuid"
)

type UserDTO struct {
	Username string `json:"username" gorm:"unique;not null"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
}

func UserDTOtoModel(request UserDTO) apps.User {
	return apps.User{
		UserId:   uuid.New().String(),
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
	}
}
