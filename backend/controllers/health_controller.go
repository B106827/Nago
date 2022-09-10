package controllers

import (
    "net/http"

    "github.com/labstack/echo/v4"
)

type HealthController struct{}

func NewHealthController() *HealthController {
    return &HealthController{}
}

// ヘルスチェック
func (hc *HealthController) Index(c echo.Context) error {
    return c.JSON(http.StatusOK, successResponse("healthy"))
}
