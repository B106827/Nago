package controllers

import (
	"NagoBackend/models"
	"NagoBackend/server/contexts"
	"net/http"

	cartForms "NagoBackend/forms/cart"

	"github.com/labstack/echo/v4"
)

type CartController struct{}

// カート情報取得
func (cac *CartController) Index(c echo.Context) error {
	user := c.Get("user").(*models.User)
	cm := models.Cart{}
	cartList, err := cm.FindByUserId(user.ID)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusOK, serverErrorResponse([]string{"エラーが発生しました"}))
	}
	if cartList == nil {
		return c.JSON(http.StatusOK, notFoundResponse(nil))
	}
	pm := models.Product{}
	userCart := []map[string]interface{}{}
	for _, cart := range cartList {
		product, err := pm.GetActiveProductWithImages(cart.ProductID)
		if err != nil || product == nil {
			continue
		}
		userCart = append(userCart, map[string]interface{}{
			"cartId":        cart.ID,
			"productId":     product.ID,
			"productName":   product.Name,
			"num":           cart.Num,
			"productPrice":  product.Price,
			"productImages": product.Images,
		})
	}
	return c.JSON(http.StatusOK, successResponse(map[string]interface{}{
		"cartList": userCart,
	}))
}

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
