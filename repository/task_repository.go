package repository

import (
	"Sistem-Manajemen-Tugas/entity"
)

type TaskRepository interface {
	FindById(string) *entity.Task
}