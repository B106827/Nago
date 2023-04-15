package controllers

import (
	"NagoBackend/config"
	"NagoBackend/database"
	"NagoBackend/handlers"
	"NagoBackend/models"
	"NagoBackend/server/contexts"
	"net/http"
	"time"

	orderCheckForms "NagoBackend/forms/order"
	orderDeliveryInfoForms "NagoBackend/forms/order"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type OrderController struct{}

func (oc *OrderController) Create(c echo.Context) error {
	orderDeliveryInfoForm := new(orderDeliveryInfoForms.OrderDeliveryInfoForm)
	cc := c.(*contexts.CustomContext)
	if err := cc.BindValidate(orderDeliveryInfoForm); err != nil {
		return c.JSON(http.StatusOK, customValidErrResponse(err))
	}
	user := c.Get("user").(*models.User)
	cm := models.Cart{}
	cartList, err := cm.FindByUserIdWithProduct(user.ID)
	if err != nil || cartList == nil {
		return c.JSON(http.StatusOK, serverErrorResponse([]string{"エラーが発生しました"}))
	}
	total := uint(0)
	for _, cart := range cartList {
		total += cart.Num * cart.Product.Price
	}
	conf := config.GetConfig()
	cartTotal := float64(total) * (1 + conf.GetFloat64("price.taxRate"))
	if uint(cartTotal) != orderDeliveryInfoForm.Total {
		return c.JSON(http.StatusOK, serverErrorResponse([]string{"エラーが発生しました"}))
	}

	// トランザクション内で処理する
	err = database.ExecuteInTx(func(tx *gorm.DB) error {
		om := models.Order{}
		om.UserID = user.ID
		om.TotalPrice = uint(cartTotal)
		om.OrderedAt = time.Now()
		if err := om.Create(tx); err != nil {
			return err
		}
		odim := models.OrderDeliveryInfo{}
		if err := odim.Create(tx, user.ID, om.ID, *orderDeliveryInfoForm); err != nil {
			return err
		}
		// Stripe Checkoutセッションを開始する
		stripeHandler := handlers.StripeHandler{}
		s, err := stripeHandler.CreateSession(om.ID, om.TotalPrice)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, successResponse(map[string]interface{}{
			"orderSession": s,
		}))
	})
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusOK, serverErrorResponse([]string{"エラーが発生しました"}))
	}
	return nil
}

func (oc *OrderController) Check(c echo.Context) error {
	orderCheckForm := new(orderCheckForms.OrderCheckForm)
	cc := c.(*contexts.CustomContext)
	if err := cc.BindValidate(orderCheckForm); err != nil {
		return c.JSON(http.StatusOK, customValidErrResponse(err))
	}

	// Stripe Checkoutセッションを確認する
	//conf := config.GetConfig()
	//stripe.Key = conf.GetString("stripe.testKey")

	//om := models.Order{}
	//order, err := om.FindById(orderCheckForm.OrderId)
	//if err != nil || order == nil {
	//	// エラーもしくはデータが存在しない
	//}
	//pp.Print(orderCheckForm)
	return nil
}
