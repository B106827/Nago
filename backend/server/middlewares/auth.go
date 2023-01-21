package middlewares

import (
	"NagoBackend/config"
	"NagoBackend/constants"
	"NagoBackend/models"
	"NagoBackend/utils/types"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type AuthMiddleware struct {
	echo.Context
}

func (am *AuthMiddleware) Authenticate() echo.MiddlewareFunc {
	// JWTの有効性確認
	conf := config.GetConfig()

	jwtConfig := middleware.JWTConfig{
		Claims:         &types.JwtCustomClaims{},
		SigningKey:     []byte(conf.GetString("session.secret")),
		TokenLookup:    "cookie:" + constants.SESSION_COOKIE_KEY_NAME,
		SuccessHandler: setUser,
	}
	return middleware.JWTWithConfig(jwtConfig)
}

func setUser(c echo.Context) {
	u := c.Get("user").(*jwt.Token)
	claims := u.Claims.(*types.JwtCustomClaims)
	userId := int(claims.UserID)
	um := models.User{}
	user, err := um.FindById(userId)
	if err != nil {
		// TODO: panic以外でメンバーが見つからない場合を実装したい
		panic("user not found")
	}
	c.Set("user", user)
}
