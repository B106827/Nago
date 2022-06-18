package controllers

import (
    "net/http"

    "github.com/labstack/echo/v4"
)

// HealthController is controller for health
type HealthController struct{}

// NewHealthController is constructor for HealthController
func NewHealthController() *HealthController {
    return new(HealthController)
}

// Index is returns http success response
func (hc *HealthController) Index(c echo.Context) error {
    return c.JSON(http.StatusOK, newResponse(
        http.StatusOK,
        http.StatusText(http.StatusOK),
        "hoge",
    ))
}
