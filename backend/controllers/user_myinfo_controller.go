package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserMyinfoController struct{}

func NewUserMyinfoController() *UserMyinfoController {
	return &UserMyinfoController{}
}

func (umc *UserMyinfoController) Index(c echo.Context) error {
	return c.JSON(http.StatusOK, successResponse(map[string]interface{}{
		"message": "success",
	}))
}
