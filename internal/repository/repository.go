package repository

import "gorm.io/gorm"

type Repository struct {
	UserRepo UserRepoItf
	TaskRepo TaskRepoItf
}

func NewRepository(db *gorm.DB) *Repository {
	userRepo := NewUserRepository(db)
	taskRepo := NewTaskRepository(db)

	return &Repository{
		UserRepo: userRepo,
		TaskRepo: taskRepo,
	}
}
