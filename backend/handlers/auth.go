package handlers

import (
	"NagoBackend/config"
	"NagoBackend/constants"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type Auth struct {
	echo.Context
}

type JwtCustomClaims struct {
	UserID uint
	jwt.StandardClaims
}

func (*Auth) Login(c echo.Context, userid uint) error {
	conf := config.GetConfig()
	claims := &JwtCustomClaims{
		userid,
		jwt.StandardClaims{
			// 1日
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(conf.GetInt("session.expireHour"))).Unix(),
		},
	}
	// create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// generate encoded token
	signingKey := []byte(conf.GetString("session.secret"))
	t, err := token.SignedString(signingKey)
	if err != nil {
		return err
	}
	// set cookie
	cookie          := new(http.Cookie)
	cookie.Expires   = time.Now().Add(24 * time.Hour) // 1日
	cookie.HttpOnly  = true
	cookie.Secure    = true
	cookie.Name      = constants.SESSION_COOKIE_KEY_NAME
	cookie.Value     = t
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
