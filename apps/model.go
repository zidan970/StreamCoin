package apps

import (
	"time"
)

type User struct {
	// gorm.Model
	UserId   string `gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique;not null"`
	Fullname string `json:"fullname"`
	Phone    string `json:"phone"`
	Email    string `json:"email" gorm:"unique;not null"`
	Address  string `json:"address"`
	Balance  uint   `json:"balance"`
	Password string `json:"password" gorm:"not null"`
}

type Login struct {
	Login_id   uint
	User_id    uint
	Login_time time.Time
}
