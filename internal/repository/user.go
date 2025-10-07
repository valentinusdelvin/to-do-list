package repository

import (
	"todolist-go/internal/domain/entity"

	"gorm.io/gorm"
)

type UserRepoItf interface {
	CreateUser(param *entity.User) error
	GetUserByEmail(email string) (*entity.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepoItf {
	return &UserRepository{db: db}
}

func (u *UserRepository) CreateUser(param *entity.User) error {
	err := u.db.Create(&param).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) GetUserByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := u.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
