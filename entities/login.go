package entities

import "time"

type Login struct {
	Login_id   uint
	User_id    uint
	Login_time time.Time
}
