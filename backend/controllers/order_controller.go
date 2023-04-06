package controllers

import (
	"NagoBackend/server/contexts"
	"net/http"

	orderAddressForms "NagoBackend/forms/order"

	"github.com/labstack/echo/v4"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/checkout/session"
)

type OrderController struct{}

func (oc *OrderController) Create(c echo.Context) error {
	orderAddressForm := new(orderAddressForms.OrderAddressForm)
	cc := c.(*contexts.CustomContext)
	if err := cc.BindValidate(orderAddressForm); err != nil {
		return c.JSON(http.StatusOK, customValidErrResponse(err))
	}
	stripe.Key = "sk_test_uSulkkzFxWPELicylSa7jMb6"
	params := &stripe.CheckoutSessionParams{
		Mode: stripe.String(string(stripe.CheckoutSessionModePayment)),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			&stripe.CheckoutSessionLineItemParams{
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					Currency: stripe.String("jpy"),
					ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
						Name: stripe.String("T-shirt"),
					},
					UnitAmount: stripe.Int64(2000),
				},
				Quantity: stripe.Int64(1),
			},
		},
		SuccessURL: stripe.String("http://localhost:3001/order/success"),
		CancelURL:  stripe.String("http://localhost:3001/order/cancel"),
	}

	s, err := session.New(params)

	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusOK, serverErrorResponse([]string{"エラーが発生しました"}))
	}
	return c.JSON(http.StatusOK, successResponse(map[string]interface{}{
		"orderSession": s,
	}))
}
