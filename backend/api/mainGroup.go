package handlers

import (
    "NagoBackend/handlers/controllers"
    "github.com/labstack/echo/v4"
)

func MainGroup(e *echo.Echo) {
    e.GET("/", controllers.Index)
}
