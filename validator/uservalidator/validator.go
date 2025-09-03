package uservalidator

import (
	"game/entity"
)

const (
	PhoneNUmberRegex = `^09[0-9]{9}$`
	PasswordRegex    = `^[A-Za-z\d!@#\$%\^&\*]{8,}$`
)

type Repository interface {
	IsUniquePhoneNumber(phoneNumber string) (bool, error)
	GetUserByPhoneNumber(phoneNumber string) (entity.User, error)
}

type Validator struct {
	repo Repository
}

func New(repository Repository) Validator {
	return Validator{repo: repository}

}
