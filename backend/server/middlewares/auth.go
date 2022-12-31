package middlewares

import (
	"NagoBackend/utils/types"
	"NagoBackend/config"
	//"NagoBackend/constants"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	//"github.com/k0kubun/pp"
)

type AuthMiddleware struct{}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (am *AuthMiddleware) Authenticate() echo.MiddlewareFunc {
	// JWTの有効性確認 
	conf := config.GetConfig()
	jwtConfig := middleware.JWTConfig{
		Claims:     &types.JwtCustomClaims{},
		SigningKey: []byte(conf.GetString("session.secret")),
		//TokenLookup: "cookie:" + constants.SESSION_COOKIE_KEY_NAME,
	}
	return middleware.JWTWithConfig(jwtConfig)
}
