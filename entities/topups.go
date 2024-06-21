package entities

import "time"

type TopUps struct {
	Topup_id     uint
	User_id      uint
	Topup_amount uint
	Topup_time   time.Time
}
