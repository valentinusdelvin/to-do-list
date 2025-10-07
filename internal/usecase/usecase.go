package usecase

import (
	"todolist-go/internal/pkg/bcrypt"
	jwt2 "todolist-go/internal/pkg/jwt"
	"todolist-go/internal/repository"
)

type Usecase struct {
	UserUsecase UserUsecaseItf
	TaskUsecase TaskUsecaseItf
}

type UsecaseDependencies struct {
	Repository *repository.Repository
	Bcrypt     bcrypt.Bcrypt
	JWT        jwt2.JWT
}

func NewUsecase(param UsecaseDependencies) *Usecase {
	UserUsecase := NewUserUsecase(param.Repository.UserRepo, param.JWT, param.Bcrypt)
	TaskUsecase := NewTaskUsecase(param.Repository.TaskRepo)

	return &Usecase{
		UserUsecase: UserUsecase,
		TaskUsecase: TaskUsecase,
	}
}
