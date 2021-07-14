package entity

type Journal struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique"`
	Type uint
	ISBN string
	CN   string
}

type JournalType struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique"`
}
