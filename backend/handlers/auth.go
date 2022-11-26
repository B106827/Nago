package handlers

import (
    "context"
    "NagoBackend/config"
    "time"

    "github.com/go-redis/redis/v8"
    "github.com/golang-jwt/jwt"
    "github.com/gorilla/sessions"
    "github.com/labstack/echo/v4"
    "github.com/rbcervilla/redisstore/v8"
    "github.com/k0kubun/pp"
)

type Auth struct {
    echo.Context
}

type jwtCustomClaims struct {
    UserID uint
    jwt.StandardClaims
}

var conf = config.GetConfig()

func (*Auth) Login(c echo.Context, userid uint) error {
    var conf = config.GetConfig()
    var test = conf.GetString("session.secret")
    pp.Print(conf)
    pp.Print(test)
    claims := &jwtCustomClaims{
        userid,
        jwt.StandardClaims{
            ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
        },
    }
    // create token with claims
    token      := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    // generate encoded token
    singingKey := []byte("qFsMnunj87WPesc31OWoPmWbBVSjqv9u")
    t, err := token.SignedString(singingKey)
    if err != nil {
        return err
    }

    session := getSession(c)
    session.Values["nago-token"] = t
    if err := session.Save(c.Request(), c.Response()); err != nil {
        c.Logger().Error(err)
        return err
    }
    return nil
}

func getSession(c echo.Context) *sessions.Session {
    client := redis.NewClient(&redis.Options{
        Addr: "nago-redis:6379",
        Password: "rIguo7jfxkMjA5Il6RyG",
    })
    store, err := redisstore.NewRedisStore(context.Background(), client)
    if err != nil {
        c.Logger().Error(err)
    }
    store.KeyPrefix("session_")
    store.Options(sessions.Options{
        MaxAge: 86400 * 1, // 1日
        HttpOnly: true,
    })
    session, err := store.Get(c.Request(), "nago-session")
    if err != nil {
        c.Logger().Error(err)
    }
    return session
}
