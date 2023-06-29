package service

import (
	"errors"

	"github.com/AvinFajarF/entity"
	"github.com/AvinFajarF/repository"
)

type TaskService struct {
	Repository repository.TaskRepository
}

func (service TaskService) GetTaskById(id string) (*entity.Task, error) {
	task := service.Repository.FindById(id)
	if task == nil {
		return task, errors.New("Task is not found")
	} else {
		return task, nil
	}
}
