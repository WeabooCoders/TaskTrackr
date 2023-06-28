package service

import (
	"Sistem-Manajemen-Tugas/entity"
	"Sistem-Manajemen-Tugas/repository"
)

type TaskService struct {
	Repository repository.TaskRepository
}

