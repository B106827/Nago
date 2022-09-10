package controllers

import (
    "github.com/labstack/echo/v4"
    "net/http"
    "NagoBackend/server/contexts"
    registerForms "NagoBackend/forms/register"
)

type RegisterController struct{}

func NewRegisterController() *RegisterController {
    return &RegisterController{}
}

// 登録URL確認
func (rc *RegisterController) URLCheck(c echo.Context) error {
    registerUrlForm := new(registerForms.RegisterUrlForm)
    cc := c.(*contexts.CustomContext)
    if err := cc.BindValidate(registerUrlForm); err != nil {
        return err
    }
    return c.JSON(http.StatusOK, successResponse("OKです"))
}

// 登録処理
func (rc *RegisterController) Register(c echo.Context) error {
    registerForm := new(registerForms.RegisterForm)
    cc := c.(*contexts.CustomContext)
    if err := cc.BindValidate(registerForm); err != nil {
        return err
    }
    return c.JSON(http.StatusOK, successResponse("OKです"))
}
