package forms

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type CartForm struct {
	ProductId uint `json:"productId"`
	CartNum   uint `json:"cartNum"`
}

func (f CartForm) Validate() error {
	return validation.ValidateStruct(&f,
		validation.Field(
			&f.ProductId,
			validation.Required.Error("パラメータが不正です"),
		),
		validation.Field(
			&f.CartNum,
			validation.Required.Error("パラメータが不正です"),
		),
	)
}
