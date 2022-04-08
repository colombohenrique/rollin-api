package models

import (
	"time"

	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	Name     string     `json:"name" validate:"nonzero"`
	Birthday *time.Time `json:"birthday" validate:"nonzero"`
}

type User struct {
	gorm.Model
	Login    string `json:"login" validate:"nonzero"`
	Password string `json:"password" validate:"nonzero"`
	Person   Person `json:"person" gorm:"embedded;embeddedPrefix:person_"`
}

type EventParticipant struct {
	gorm.Model
	Participant User `json:"participant" gorm:"embedded;embeddedPrefix:user_"`
	Event_id    int64
	Event       Event `json:"event" gorm:"embedded;embeddedPrefix:event_;foreignkey:Event_id; references:ID"`
}

func DataValidate(user *User) error {
	if err := validator.Validate(user); err != nil {
		return err
	}
	return nil
}
