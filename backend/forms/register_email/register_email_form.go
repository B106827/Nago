package forms

type RegisterEmailForm struct {
    Email        string `json:"email" validate:"required,email,max=255"`
    RedirectPath string `json:"redirectPath" validate:"max=255"`
}
