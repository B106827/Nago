package forms

import (
    validation "github.com/go-ozzo/ozzo-validation/v4"
    "github.com/go-ozzo/ozzo-validation/v4/is"
)

type RegisterEmailForm struct {
    Email string `json:"email"`
}

func (f RegisterEmailForm) Validate() error {
    return validation.ValidateStruct(&f,
        validation.Field(
            &f.Email,
            validation.Required.Error("メールアドレスは入力必須です"),
            validation.RuneLength(5, 255).Error("メールアドレスは{min}〜{max}文字です"),
            is.Email.Error("メールアドレスが不正な形跡です"),
        ),
    )
}
