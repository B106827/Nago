package middlewares

import (
	"github.com/labstack/echo/v4"
)

type AuthMiddleware struct{}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (am *AuthMiddleware) Authenticate(next echo.HandlerFunc) error {
	return nil
}
