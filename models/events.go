package models

import (
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Name         string             `json:"name"`
	Creator      User               `json:"creator" gorm:"embedded;embeddedPrefix:creator_"`
	Participants []EventParticipant `json:"participants" gorm:"-:all"`
	Items        []Item             `json:"items" gorm:"-:all"`
}
