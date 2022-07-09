package forms

type RegisterEmailForm struct {
    Email    string `json:"email" validate:"required,email,max=255"`
    Password string `json:"password" validate:"required"`
}
