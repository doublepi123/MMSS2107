package entity

import "time"

type Userid struct {
	Username  string `gorm:"primarykey"`
	Userkey   string
	Validtime time.Time
}
