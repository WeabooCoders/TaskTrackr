package repository

import (
	"github.com/AvinFajarF/entity"
)

type TaskRepository interface {
	FindById(string) *entity.Task
}
