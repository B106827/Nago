package handlers

import (
    "context"
    "fmt"
    "NagoBackend/config"
    "NagoBackend/constants"
    "time"

    "github.com/go-redis/redis/v8"
    "github.com/golang-jwt/jwt"
    "github.com/gorilla/sessions"
    "github.com/labstack/echo/v4"
    "github.com/rbcervilla/redisstore/v8"
)

type Auth struct {
    echo.Context
}

type jwtCustomClaims struct {
    UserID uint
    jwt.StandardClaims
}

func (*Auth) Login(c echo.Context, userid uint) error {
    conf := config.GetConfig()
    claims := &jwtCustomClaims{
        userid,
        jwt.StandardClaims{
            // 1日
            ExpiresAt: time.Now().Add(time.Hour * time.Duration(conf.GetInt("session.expireHour"))).Unix(),
        },
    }
    // create token with claims
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    // generate encoded token
    singingKey := []byte(conf.GetString("session.secret"))
    t, err := token.SignedString(singingKey)
    if err != nil {
        return err
    }

    session := getSession(c)
    session.Values[constants.SESSION_KEY_NAME] = t
    if err := session.Save(c.Request(), c.Response()); err != nil {
        c.Logger().Error(err)
        return err
    }
    return nil
}

func getSession(c echo.Context) *sessions.Session {
    conf := config.GetConfig()
    client := redis.NewClient(&redis.Options{
        Addr: fmt.Sprintf("%s:%d", conf.GetString("redis.host"), conf.GetInt("redis.port")),
        Password: conf.GetString("redis.password"),
    })
    store, err := redisstore.NewRedisStore(context.Background(), client)
    if err != nil {
        c.Logger().Error(err)
    }
    store.KeyPrefix(constants.SESSION_REDIS_KEY_PREFIX)
    store.Options(sessions.Options{
        MaxAge: 3600 * conf.GetInt("session.expireHour"), // 1日
        HttpOnly: true,
        Secure: true,
    })
    session, err := store.Get(c.Request(), constants.SESSION_COOKIE_KEY_NAME)
    if err != nil {
        c.Logger().Error(err)
    }
    return session
}
