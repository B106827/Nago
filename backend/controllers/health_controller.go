package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HealthController struct{}

// ヘルスチェック
func (hc *HealthController) Index(c echo.Context) error {
	return c.JSON(http.StatusOK, successResponse("healthy"))
}
