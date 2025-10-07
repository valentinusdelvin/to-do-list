package repository

import "gorm.io/gorm"

type TaskRepoItf interface {
}

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepoItf {
	return &TaskRepository{db: db}
}
