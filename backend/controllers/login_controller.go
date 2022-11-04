package controllers

import (
    "NagoBackend/server/contexts"
    "NagoBackend/models"
    "net/http"

    loginForms "NagoBackend/forms/login"
    "github.com/labstack/echo/v4"

//    "github.com/k0kubun/pp"
)

type LoginController struct{}

func NewLoginController() *LoginController {
    return &LoginController{}
}

// ログイン処理
func (lc *LoginController) Login(c echo.Context) error {
    loginForm := new(loginForms.LoginForm)
    cc := c.(*contexts.CustomContext)
    if err := cc.BindValidate(loginForm); err != nil {
        return c.JSON(http.StatusOK, badRequestResponse(err))
    }
    user := new(models.User)
    if err := user.FindByEmail(loginForm.Email); err != nil {
        return c.JSON(http.StatusOK, badRequestResponse([]string{"メンバーが見つかりません"}))
    }
    return c.JSON(http.StatusOK, successResponse(map[string]interface{}{
        "user": user,
        "message": "ログインしました",
    }))
}

// ログアウト処理
func (lc *LoginController) Logout(c echo.Context) error {
    return c.JSON(http.StatusOK, successResponse(map[string]interface{}{
        "message": "ログアウトしました",
    }))
}
