package services

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

func NewToken(userId uint) (*string, time.Time, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(userId)),
		ExpiresAt: oneDay().Unix(),
	})

	key := os.Getenv("JWT_KEY")
	token, err := claims.SignedString([]byte(key))
	if err != nil {
		return nil, time.Time{}, err
	}

	return &token, oneDay(), nil
}

func ParseJwt(cookie string) (*jwt.StandardClaims, error) {
	key := os.Getenv("JWT_KEY")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if err != nil {
		return nil, err
	}

	claims := token.Claims.(*jwt.StandardClaims)

	return claims, nil
}

func oneDay() time.Time {
	return time.Now().Add(time.Hour * 24)
}
