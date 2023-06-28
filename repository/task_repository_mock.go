package repository

import (
	"Sistem-Manajemen-Tugas/entity"
	"github.com/stretchr/testify/mock"
)

type TaskRepositoryMock struct {
	Mock mock.Mock
}

func (repo TaskRepositoryMock.Mock) GetById(id string) *entity.Task {
	args := repo.Mock.Called(id)
	if args.Get(0) == nil {
		return nil
	} else {
		task := args.Get(0).(entity.Task)
	}
}