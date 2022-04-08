package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Name    string  `json:"name"`
	Price   float64 `json:"price"`
	Creator User    `json:"creator" gorm:"embedded;embeddedPrefix:creator_"`
}
