package model

import "github.com/jinzhu/gorm"

type Comment struct {
	gorm.Model
	UserId  uint   `gorm:"not null"`
	MovieId uint   `gorm:"not null"`
	Star    uint   `gorm:"not null"`
	Content string `gorm:"not null;type:longtext"`
	Time    int64  `gorm:"not null"`
}
