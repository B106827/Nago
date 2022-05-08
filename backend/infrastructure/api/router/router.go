package router

import (
    "github.com/labstack/echo/v4"
    "NagoBackend/infrastructure/api/handler"
)

func NewRouter(e *echo.Echo, handler handler.AppHandler) {
    e.GET("/users", handler.GetUsers)
}
