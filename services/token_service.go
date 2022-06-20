package services

import (
	"jewete/entities"
	"os"

	"github.com/golang-jwt/jwt"
)

func NewToken(claims *entities.JWTClaim) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	key := os.Getenv("JWT_KEY")
	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}

func ParseJwt(cookie *string) (*jwt.StandardClaims, error) {
	key := os.Getenv("JWT_KEY")

	token, err := jwt.ParseWithClaims(*cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if err != nil {
		return nil, err
	}

	claims := token.Claims.(*jwt.StandardClaims)
	return claims, nil
}
