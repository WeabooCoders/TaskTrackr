package repository

import (
	"github.com/AvinFajarF/entity"

	"github.com/stretchr/testify/mock"
)

type TaskRepositoryMock struct {
	Mock mock.Mock
}

func (repository TaskRepositoryMock) FindById(id string) *entity.Task {
	args := repository.Mock.Called(id)
	if args.Get(0) == nil {
		return nil
	} else {
		task := args.Get(0).(entity.Task)
		return &task
	}
}
