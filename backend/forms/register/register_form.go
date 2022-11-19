package forms

import (
    validation "github.com/go-ozzo/ozzo-validation/v4"
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
