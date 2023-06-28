import (
	"Sistem-Manajemen-Tugas/entity"
	"Sistem-Manajemen-Tugas/repository"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

package service

var taskRepository = &repository.TaskRepositoryMock{Mock: mock.Mock{}}
var taskService = TaskService{Repository: taskRepository}

func TestService_GetById(t *testing.T) {
// Todo: Add test
}