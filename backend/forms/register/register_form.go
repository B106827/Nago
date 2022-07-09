package forms

type RegisterUrlForm struct {
    ID string `json:"id" validate:"required"`
}

type RegisterForm struct {
    TmpID    string `json:"tmpId" validate:"required"`
    Name     string `json:"name" validate:"required"`
    Email    string `json:"email" validate:"required,email,max=255"`
    Password string `json:"password" validate:"required,min=6,max=20"`
}
