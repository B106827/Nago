package forms

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type RegisterUrlForm struct {
	TmpID string `json:"tmpId"`
}

func (f RegisterUrlForm) Validate() error {
	return validation.ValidateStruct(&f,
		validation.Field(
			&f.TmpID,
			validation.Required.Error("無効なURLです"),
		),
	)
}

type RegisterForm struct {
	TmpID    string `json:"tmpId"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (f RegisterForm) Validate() error {
	return validation.ValidateStruct(&f,
		validation.Field(
			&f.TmpID,
			validation.Required.Error("無効なURLです"),
		),
		validation.Field(
			&f.Name,
			validation.Required.Error("お名前は入力必須です"),
			validation.RuneLength(1, 255).Error("お名前は {min}〜{max} 文字で入力してください"),
		),
		validation.Field(
			&f.Email,
			validation.Required.Error("メールアドレスは入力必須です"),
			validation.RuneLength(5, 255).Error("メールアドレスは {min}〜{max} 文字で入力してください"),
			is.Email.Error("メールアドレスが不正な形式です"),
		),
		validation.Field(
			&f.Password,
			validation.Required.Error("パスワードは入力必須です"),
			validation.RuneLength(6, 20).Error("パスワードは {min}〜{max} 文字で入力してください"),
		),
	)
}
