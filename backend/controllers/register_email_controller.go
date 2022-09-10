package controllers

import (
    "github.com/labstack/echo/v4"
    "net/http"
    "NagoBackend/server/contexts"
    registerEmailForms "NagoBackend/forms/register_email"
)

type RegisterEmailController struct{}

func NewRegisterEmailController() *RegisterEmailController {
    return &RegisterEmailController{}
}

// メールアドレス登録処理
func (rec *RegisterEmailController) RegisterEmail(c echo.Context) (err error) {
    registerEmailForm := new(registerEmailForms.RegisterEmailForm)
    cc := c.(*contexts.CustomContext)
    if err := cc.BindValidate(registerEmailForm); err != nil {
        return err
    }
    return c.JSON(http.StatusOK, successResponse("OKです"))
}
