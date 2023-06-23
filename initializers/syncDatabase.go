package initializers

import (
	"github.com/AvinFajarF/model"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&model.User{}, &model.Task{})
}
