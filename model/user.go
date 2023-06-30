package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Uuid        string `gorm:"primaryKey"`
	Username  string `gorm:"uniqueIndex"`
	Password  string
	Email     string `gorm:"uniqueIndex"`
	Tasks     []Task
	CreatedAt time.Time
	UpdatedAt time.Time
}
