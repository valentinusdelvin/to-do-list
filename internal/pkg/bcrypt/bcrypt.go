package bcrypt

import (
	"golang.org/x/crypto/bcrypt"
)

type BcryptItf interface {
	HashPassword(password string) (string, error)
	ComparePassword(hashedPassword, password string) error
}

type Bcrypt struct {
	cost int
}

func NewBcrypt(cost int) *Bcrypt {
	return &Bcrypt{
		cost: cost,
	}
}

func (b *Bcrypt) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), b.cost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (b *Bcrypt) ComparePassword(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
