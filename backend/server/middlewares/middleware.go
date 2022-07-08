package middlewares

import (
    "github.com/go-playground/validator"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "NagoBackend/config"
    "net/http"
)

type CustomValidator struct {
    validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
    return cv.validator.Struct(i)
}

func InitMiddleware(e *echo.Echo) {
    c := config.GetConfig()
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())
    e.Use(middleware.CORSWithConfig(middleware.CORSConfig {
        AllowOrigins: c.GetStringSlice("server.cors"),
        AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
    }))
    e.Validator = &CustomValidator{validator: validator.New()}
}
