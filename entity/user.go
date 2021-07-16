package entity

type User struct {
	Username string `gorm:"primarykey"`
	Password string
}
