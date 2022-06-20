package entities

import (
	"github.com/golang-jwt/jwt"
)

type JWTClaim struct {
	User *User
	*jwt.StandardClaims
}
