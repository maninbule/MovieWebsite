package model

import (
	"github.com/jinzhu/gorm"
)

type Movie struct {
	gorm.Model
	Title       string `gorm:"index;not null"`
	Actor       string `gorm:"not null"`
	Description string `gorm:"not null"`
	MovieTime   int64  `gorm:"not null"`
	PostImage   string `gorm:"not null"`
}
