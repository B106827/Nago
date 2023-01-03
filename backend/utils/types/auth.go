package types

import (
	"github.com/golang-jwt/jwt"
)

type JwtCustomClaims struct {
	UserID uint
	jwt.StandardClaims
}
