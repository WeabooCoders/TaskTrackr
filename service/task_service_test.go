package service

import (
	"testing"

	"github.com/AvinFajarF/repository"

	"github.com/stretchr/testify/mock"
)

var taskRepository = &repository.TaskRepositoryMock{Mock: mock.Mock{}}
var taskService = TaskService{Repository: taskRepository}

func TestService_GetById(t *testing.T) {
	// Todo: Add test
}
