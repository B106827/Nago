package controllers

import (
	"NagoBackend/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type MasterController struct{}

// 商品一覧
func (pc *MasterController) GetPrefMaster(c echo.Context) error {
	pmm := models.PrefMaster{}
	prefMaster, err := pmm.GetPrefMaster()
	if err != nil || prefMaster == nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusOK, serverErrorResponse([]string{"エラーが発生しました"}))
	}
	return c.JSON(http.StatusOK, successResponse(map[string]interface{}{
		"prefMaster": prefMaster,
	}))
}
