package handlers

import (
	"NagoBackend/config"
	"NagoBackend/constants"
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/sessions"
	"github.com/k0kubun/pp"
	"github.com/labstack/echo/v4"
	"github.com/rbcervilla/redisstore/v8"
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
	pp.Print(t)

	session := getSession(c)
	// セッション保存
	session.Values[constants.SESSION_KEY_NAME] = t
	session.Values["test"] = "hoge"
	if err := session.Save(c.Request(), c.Response()); err != nil {
		c.Logger().Error(err)
		return err
	}
	return nil
}

func (*Auth) Logout(c echo.Context) error {
	session := getSession(c)
	// セッション削除
	session.Options.MaxAge = -1
	if err := session.Save(c.Request(), c.Response()); err != nil {
		c.Logger().Error(err)
		return err
	}
	return nil
}

func (*Auth) Test(c echo.Context) error {
	session := getSession(c)
	pp.Print(session)
	return nil
}

func getSession(c echo.Context) *sessions.Session {
	conf := config.GetConfig()
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.GetString("redis.host"), conf.GetInt("redis.port")),
		Password: conf.GetString("redis.password"),
	})
	store, err := redisstore.NewRedisStore(context.Background(), client)
	if err != nil {
		c.Logger().Error(err)
	}
	store.KeyPrefix(constants.SESSION_REDIS_KEY_PREFIX)
	store.Options(sessions.Options{
		MaxAge:   3600 * conf.GetInt("session.expireHour"), // 1日
		Secure:   true,
		HttpOnly: true,
	})
	session, err := store.Get(c.Request(), constants.SESSION_COOKIE_KEY_NAME)
	if err != nil {
		c.Logger().Error(err)
	}
	return session
}
