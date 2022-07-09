package controllers

import (
    "github.com/labstack/echo/v4"
    "net/http"
    loginForms "NagoBackend/forms/login"
    "NagoBackend/server/contexts"
)

type LoginController struct{}

func NewLoginController() *LoginController {
    return new(LoginController)
}

// ログイン処理
func (lc *LoginController) Login(c echo.Context) error {
    loginForm := new(loginForms.LoginForm)
    cc := c.(*contexts.CustomContext)
    if err := cc.BindValidate(loginForm); err != nil {
        return err
    }
    return c.JSON(http.StatusOK, successResponse("OKです"))
}

// ログアウト処理
func (lc *LoginController) Logout(c echo.Context) error {
    return c.JSON(http.StatusOK, successResponse("OKです"))
}
