package controllers

import (
    "github.com/labstack/echo/v4"
    "net/http"
    loginForms "NagoBackend/forms/login"
    "NagoBackend/server/contexts"
    "NagoBackend/models"
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
    user := new(models.User)
    if err := user.FindByEmail(loginForm.Email); err != nil {
        return c.JSON(http.StatusOK, successResponse("メンバーが見つかりません"))
    }
    return c.JSON(http.StatusOK, user.Name)
}

// ログアウト処理
func (lc *LoginController) Logout(c echo.Context) error {
    return c.JSON(http.StatusOK, successResponse("OKです"))
}
