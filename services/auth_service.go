package services

import (
	"jewete/entities"
	"jewete/repositories"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var authRepository = repositories.NewAuth()

func CreateUser(user *entities.User) (*entities.User, error) {
	user.Password = hashPassword(string(user.Password))
	err := authRepository.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetUser(user *entities.User) (*entities.User, error) {
	record, _ := authRepository.Get(user)
	if record.ID == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	if isCorrectPassword(record.Password, user.Password) {
		return record, nil
	}

	return nil, gorm.ErrRecordNotFound
}

func isCorrectPassword(hashed string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	return err == nil
}

func hashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		panic("hash password error : " + err.Error())
	}

	return string(hashedPassword)
}
