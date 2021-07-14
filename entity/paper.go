package entity

import "time"

type Paper struct {
	ID            uint   `gorm:"primaryKey"`
	PaperID       string `gorm:"unique"`
	Name          string `gorm:"index"`
	CName         string	`gorm:"index"`
	Type          uint8	`gorm:"index"`
	Journal       uint `gorm:"index"`
	PublishedTime time.Time `gorm:"index"`
	Diff          string
	StartPage     uint
	EndPage       uint
	NumberOfPaper uint
	Author        string
	File          string
}
