package model

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title       string
	Description string
	Status      string
	UserID      uint `gorm:"foreignKey:UserID"`
	User        User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"`
}
