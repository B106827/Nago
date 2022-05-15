package controllers

import (
    "net/http"

    "github.com/labstack/echo/v4"
)

type HealthController struct{}

func NewHealthController() *HealthController {
    return new(HealthController)
}

func (hc *HealthController) Index(c echo.Context) error {
    return c.JSON(http.StatusOK, newResponse(
        http.StatusOK,
        http.StatusText(http.StatusOK),
        "hoge",
    ))
}
