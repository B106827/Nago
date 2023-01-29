package forms

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ProductForm struct {
	Id uint `json:"productId" param:"productId"`
}

func (f ProductForm) Validate() error {
	return validation.ValidateStruct(&f,
		validation.Field(
			&f.Id,
			validation.Required.Error("無効なURLです"),
		),
	)
}
