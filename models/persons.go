package models

import (
	"time"

	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	Name     string     `json:"name" validate:"nonzero"`
	Birthday *time.Time `json:"birthday"`
}

type User struct {
	gorm.Model
	Login    string `json:"login" validate:"nonzero"`
	Password string `json:"password" validate:"nonzero"`
	Person   Person `json:"person" gorm:"embedded;embeddedPrefix:person_"`
}

func DataValidate(user *User) {

}
