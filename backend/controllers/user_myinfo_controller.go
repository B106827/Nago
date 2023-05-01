package controllers

import (
	"net/http"

	"github.com/k0kubun/pp"
	"github.com/labstack/echo/v4"
)

type UserMyinfoController struct{}

func (umc *UserMyinfoController) Index(c echo.Context) error {
	user := c.Get("user")
	pp.Print(user)
	return c.JSON(http.StatusOK, successResponse(map[string]interface{}{
		"message": "準備中です",
	}))
}
