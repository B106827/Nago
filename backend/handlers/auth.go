package handlers

import (
	"NagoBackend/config"
	"NagoBackend/constants"
	"NagoBackend/utils/types"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type Auth struct {
	echo.Context
}

func (*Auth) Login(c echo.Context, userid uint) error {
	conf := config.GetConfig()
	claims := types.JwtCustomClaims{
		userid,
		jwt.StandardClaims{
			// 1日
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(conf.GetInt("jwt.expireHour"))).Unix(),
		},
	}
	// create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// generate encoded token
	signingKey := []byte(conf.GetString("jwt.secret"))
	t, err := token.SignedString(signingKey)
	if err != nil {
		return err
	}
	// set cookie
	cookie := &http.Cookie{
		Expires:  time.Now().Add(time.Hour * time.Duration(conf.GetInt("jwt.expireHour"))),
		HttpOnly: true,
		Secure:   true,
		Name:     constants.SESSION_COOKIE_KEY_NAME,
		Value:    t,
	}
	c.SetCookie(cookie)
	return nil
}

func (*Auth) Logout(c echo.Context) error {
	cookie, err := c.Cookie(constants.SESSION_COOKIE_KEY_NAME)
	if err != nil {
		c.Logger().Error(err)
		return err
	}
	cookie.MaxAge = -1
	c.SetCookie(cookie)
	return nil
}
