package entity

type Teacher struct {
	ID        uint	`gorm:"primaryKey"`
	Name      string `gorm:"index"`
	WorkID    string `gorm:"unique"`
	Unit      string
	Phone     string
	Email     string
	Telephone string
	QQ        string
}
