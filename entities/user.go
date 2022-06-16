package entities

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `json:"username" binding:"required" gorm:"unique"`
	Email     string `json:"email" binding:"required,email" gorm:"unique"`
	Password  []byte `json:"-" binding:"required,max=15,alphanum"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (u *User) HashPassword() {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	if err != nil {
		panic("hash password error : " + err.Error())
	}

	u.Password = hashedPassword
}
