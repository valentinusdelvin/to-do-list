package usecase

import "todolist-go/internal/repository"

type TaskUsecaseItf interface {
}

type TaskUsecase struct {
	taskRepo repository.TaskRepoItf
}

func NewTaskUsecase(repo repository.TaskRepoItf) TaskUsecaseItf {
	return &TaskUsecase{taskRepo: repo}
}
