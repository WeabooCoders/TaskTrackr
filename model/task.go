package model

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Uuid          string `gorm:"primaryKey"`
	Title       string
	Description string
	Status      string `gorm:"type:ENUM('done', 'progress')"`
	UserID      uint `gorm:"foreignKey:UserID"`
	User        User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"`
}
