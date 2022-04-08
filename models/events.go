package models

import "gorm.io/gorm"

type Event struct {
	gorm.Model
	Name         string `json:"name"`
	Creator      User   `json:"creator"`
	Participants []User `json:"participants"`
	Items        []Item
}
