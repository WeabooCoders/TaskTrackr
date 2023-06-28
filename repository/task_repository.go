package repository

import (
	"Sistem-Manajemen-Tugas/entity"
)

type TaskRepository interface {
	GetById(string) *entity.Task
}