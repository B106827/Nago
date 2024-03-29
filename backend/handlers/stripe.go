package handlers

import (
	"NagoBackend/config"
	"NagoBackend/constants"
	"fmt"

	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/checkout/session"
)

type StripeHandler struct{}

// Stripe Checkout セッションを開始する
func (*StripeHandler) CreateSession(orderId uint, price uint) (*stripe.CheckoutSession, error) {
	conf := config.GetConfig()
	stripe.Key = conf.GetString("stripe.testKey")
	orderSuccessURL := fmt.Sprintf(conf.GetString("url.front")+constants.ORDER_SUCCESS_PATH+"?session_id={CHECKOUT_SESSION_ID}&order_id=%d", orderId)
	orderCancelURL  := fmt.Sprintf(conf.GetString("url.front")+constants.ORDER_CANCEL_PATH+"?order_id=%d", orderId)

	params := &stripe.CheckoutSessionParams{
		Mode: stripe.String(string(stripe.CheckoutSessionModePayment)),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			&stripe.CheckoutSessionLineItemParams{
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					Currency: stripe.String(constants.ORDER_CURRENCY),
					ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
						Name: stripe.String(constants.SERVICE_NAME),
					},
					UnitAmount: stripe.Int64(int64(price)),
				},
				Quantity: stripe.Int64(1),
			},
		},
		SuccessURL: stripe.String(orderSuccessURL),
		CancelURL:  stripe.String(orderCancelURL),
	}
	return session.New(params)
}

// Stripe Checkout データを取得する
func (*StripeHandler) CheckSessionResult(sessionId string) (*stripe.CheckoutSession, error) {
	conf := config.GetConfig()
	stripe.Key = conf.GetString("stripe.testKey")

	return session.Get(sessionId, nil)
}
