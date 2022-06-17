package services

import (
	"jewete/entities"
	"jewete/repositories"

	"golang.org/x/crypto/bcrypt"
)

var authRepository = repositories.NewAuth()

func CreateUser(user *entities.User) (*entities.User, error) {
	user.Password = HashPassword(string(user.Password))
	err := authRepository.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func HashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		panic("hash password error : " + err.Error())
	}

	return string(hashedPassword)
}
