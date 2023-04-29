package forms

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type OrderCancelForm struct {
	OrderId uint `json:"orderId"`
}

func (f OrderCancelForm) Validate() error {
	return validation.ValidateStruct(&f,
		validation.Field(
			&f.OrderId,
			validation.Required.Error("パラメータが不正です"),
		),
	)
}
