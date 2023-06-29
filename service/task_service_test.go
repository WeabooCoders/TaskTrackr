package service

import (
	"testing"

	"github.com/AvinFajarF/entity"
	"github.com/AvinFajarF/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var taskRepository = &repository.TaskRepositoryMock{Mock: mock.Mock{}}
var taskService = TaskService{Repository: taskRepository}

func TestService_GetById(t *testing.T) {
	// Mock Data atau biasa disebut data tidak real. Digunakan seolah data berasal dari database
	filteredTask := entity.Task{
		Title:       "PR Matematika",
		Description: "Soal Uji Kompetensi 1 - 10",
		Status:      "Belum dikerjakan",
		UserID:      "5",
	}
	taskRepository.Mock.On("FindById", "5").Return(filteredTask)
	task, err := taskService.GetTaskById("5")

	assert.NotNil(t, task)
	assert.Nil(t, err)
}
