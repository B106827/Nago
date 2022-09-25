package middlewares

import (
    "NagoBackend/config"
    "net/http"

    validation "github.com/go-ozzo/ozzo-validation/v4"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
)

type CustomValidator struct {}

func (cv *CustomValidator) Validate(i interface{}) error {
    if c, ok := i.(validation.Validatable); ok {
        return c.Validate()
    }
    return nil
}

func InitMiddleware(e *echo.Echo) {
    c := config.GetConfig()
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())
    e.Use(middleware.CORSWithConfig(middleware.CORSConfig {
        AllowOrigins: c.GetStringSlice("server.cors"),
        AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
    }))
    e.Validator = &CustomValidator{}
}
