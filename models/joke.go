package models

import "github.com/jinzhu/gorm"

type Joke struct {
	gorm.Model
	Value string `gorm:"not null"`
}
