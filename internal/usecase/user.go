package usecase

import (
	"errors"
	"strings"
	"todolist-go/internal/domain/dto"
	"todolist-go/internal/domain/entity"
	"todolist-go/internal/pkg/bcrypt"
	"todolist-go/internal/pkg/jwt"
	"todolist-go/internal/repository"

	"github.com/google/uuid"
)

type UserUsecaseItf interface {
	Register(param dto.Register) error
	Login(param dto.Login) (string, error)
}

type UserUsecase struct {
	userRepo repository.UserRepoItf
	jwt      jwt.JWT
	bcrypt   bcrypt.Bcrypt
}

func NewUserUsecase(repo repository.UserRepoItf, jwt jwt.JWT, bcrypt bcrypt.Bcrypt) UserUsecaseItf {
	return &UserUsecase{
		userRepo: repo,
		jwt:      jwt,
		bcrypt:   bcrypt,
	}
}

func (h *UserUsecase) Register(param dto.Register) error {
	userExist, err := h.userRepo.GetUserByEmail(param.Email)
	if err == nil && userExist != nil {
		return errors.New("user with this email already exists")
	}

	hashedPassword, err := h.bcrypt.HashPassword(param.Password)
	if err != nil {
		return err
	}

	var role string
	if strings.Contains(param.Email, "student.ub.ac.id") {
		role = "admin"
	} else {
		role = "user"
	}

	user := entity.User{
		UserId:   uuid.NewString(),
		Name:     param.Name,
		Email:    param.Email,
		Role:     role,
		Password: hashedPassword,
	}

	err = h.userRepo.CreateUser(&user)
	if err != nil {
		return err
	}

	return nil
}

func (h *UserUsecase) Login(param dto.Login) (string, error) {
	user, err := h.userRepo.GetUserByEmail(param.Email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	err = h.bcrypt.ComparePassword(user.Password, param.Password)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	token, err := h.jwt.GenerateToken(user.UserId, user.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}
