package initializers

import (
	"github.com/AvinFajarF/model"
)

func Migrate() {
	DB.AutoMigrate(&model.User{}, &model.Task{})
}
