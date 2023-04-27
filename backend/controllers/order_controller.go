package controllers

import (
	"NagoBackend/config"
	"NagoBackend/database"
	"NagoBackend/handlers"
	"NagoBackend/models"
	"NagoBackend/server/contexts"
	"encoding/json"
	"net/http"
	"time"

	orderCheckForms "NagoBackend/forms/order"
	orderDeliveryInfoForms "NagoBackend/forms/order"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"gorm.io/datatypes"
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
	// 決済詳細情報用にデータを準備しておく
	orderDetails := []map[string]uint{}
	for _, cart := range cartList {
		total += cart.Num * cart.Product.Price
		orderDetails = append(orderDetails, map[string]uint{
			"productId": cart.ProductID,
			"price":     cart.Product.Price,
			"num":       cart.Num,
		})
	}
	conf := config.GetConfig()
	cartTotal := float64(total) * (1 + conf.GetFloat64("price.taxRate"))
	if uint(cartTotal) != orderDeliveryInfoForm.Total {
		// 合計金額の確認
		return c.JSON(http.StatusOK, serverErrorResponse([]string{"エラーが発生しました"}))
	}

	// トランザクション内で処理する
	err = database.ExecuteInTx(func(tx *gorm.DB) error {
		// 決済情報の登録
		om            := models.Order{}
		om.UserID      = user.ID
		om.TotalPrice  = uint(cartTotal)
		om.OrderedAt   = time.Now()
		if err := om.CreateInTx(tx); err != nil {
			return err
		}
		// 決済詳細情報の登録
		odm := models.OrderDetail{}
		if err := odm.CreateInTx(tx, om.ID, &orderDetails); err != nil {
			return err
		}
		// 送り先情報の登録
		odim := models.OrderDeliveryInfo{}
		if err := odim.CreateInTx(tx, om.ID, *orderDeliveryInfoForm); err != nil {
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
	stripeHandler := handlers.StripeHandler{}
	s, err        := stripeHandler.CheckSessionResult(orderCheckForm.SessionId)
	if err != nil {
		// Stripe Checkoutの全体的なエラー（セッションIDが存在しない場合も含む）
		c.Logger().Error(err)
		return c.JSON(http.StatusOK, serverErrorResponse([]string{"決済情報の問い合わせでエラーが発生しました"}))
	}
	// NOTE: 現状はCheckoutセッションの結果がある段階でStripeの支払いは成功とする
	user       := c.Get("user").(*models.User)
	om         := models.Order{}
	order, err := om.FindById(orderCheckForm.OrderId)
	if err != nil || order == nil {
		return c.JSON(http.StatusOK, serverErrorResponse([]string{"決済記録が見つかりません"}))
	}
	if order.Status != models.ORDER_STATUS_PENDING || order.UserID != user.ID || s.AmountTotal != int64(order.TotalPrice) {
		return c.JSON(http.StatusOK, serverErrorResponse([]string{"決済情報が不正です"}))
	}

	sJson, err := json.Marshal(s)
	if err != nil {
		return c.JSON(http.StatusOK, serverErrorResponse([]string{"エラーが発生しました"}))
	}
	// トランザクション内で処理する
	err = database.ExecuteInTx(func(tx *gorm.DB) error {
		osdm            := models.OrderStripeDetail{}
		osdm.OrderID     = order.ID
		osdm.StripeID    = s.ID
		osdm.ResultData  = datatypes.JSON(sJson)
		if err := osdm.CreateInTx(tx); err != nil {
			return err
		}
		if err := order.UpdateStatusCompleteInTx(tx); err != nil {
			return err
		}
		cm := models.Cart{}
		if err := cm.DeleteCartDataInTx(tx, user.ID); err != nil {
			return err
		}
		return c.JSON(http.StatusOK, successResponse(map[string]interface{}{
			"message": "お支払いが正常に完了しました",
		}))
	})
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusOK, serverErrorResponse([]string{"決済に失敗しました。管理者に問い合わせてください"}))
	}
	return nil
}
