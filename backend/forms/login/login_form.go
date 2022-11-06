package forms

import (
    validation "github.com/go-ozzo/ozzo-validation/v4"
    "github.com/go-ozzo/ozzo-validation/v4/is"
)

type LoginForm struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

func (f LoginForm) Validate() error {
    return validation.ValidateStruct(&f,
        validation.Field(
            &f.Email,
            validation.Required.Error("メールアドレスは入力必須です"),
            validation.RuneLength(5, 255).Error("メールアドレスは{min}〜{max}文字です"),
            is.Email.Error("メールアドレスが不正な形式です"),
        ),
        validation.Field(
            &f.Password,
            validation.Required.Error("パスワードは入力必須です"),
            validation.RuneLength(6, 20).Error("パスワードは{min}〜{max}文字です"),
        ),
    )
}
