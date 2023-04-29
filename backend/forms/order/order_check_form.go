package forms

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type OrderCheckForm struct {
	SessionId string `json:"sessionId"`
	OrderId   uint   `json:"orderId"`
}

func (f OrderCheckForm) Validate() error {
	return validation.ValidateStruct(&f,
		validation.Field(
			&f.SessionId,
			validation.Required.Error("パラメータが不正です"),
		),
		validation.Field(
			&f.OrderId,
			validation.Required.Error("パラメータが不正です"),
		),
	)
}
