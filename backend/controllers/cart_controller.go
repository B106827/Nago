package controllers

import (
	"NagoBackend/models"
	"NagoBackend/server/contexts"
	"net/http"

	cartForms "NagoBackend/forms/cart"

	"github.com/labstack/echo/v4"
)

type CartController struct{}

func (cac *CartController) Update(c echo.Context) error {
	cartForm := new(cartForms.CartForm)
	cc := c.(*contexts.CustomContext)
	if err := cc.BindValidate(cartForm); err != nil {
		return c.JSON(http.StatusOK, badRequestResponse(err))
	}
	user := c.Get("user").(*models.User)
	cm := models.Cart{}
	cart, err := cm.FindByUserIdAndProductId(user.ID, cartForm.ProductId)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusOK, serverErrorResponse([]string{"エラーが発生しました"}))
	}
	if cart == nil {
		cm.UserID = user.ID
		cm.ProductID = cartForm.ProductId
		cm.Num = cartForm.CartNum
		if err := cm.Create(); err != nil {
			c.Logger().Error(err)
			return c.JSON(http.StatusOK, serverErrorResponse([]string{"エラーが発生しました"}))
		}
	} else {
		if err := cart.Update(cartForm.CartNum); err != nil {
			c.Logger().Error(err)
			return c.JSON(http.StatusOK, serverErrorResponse([]string{"エラーが発生しました"}))
		}
	}
	return c.JSON(http.StatusOK, successResponse(map[string]interface{}{
		"message": "カートを更新しました",
	}))
}
