package controllers

import (
    "NagoBackend/handlers"
    "NagoBackend/models"
    "NagoBackend/server/contexts"
    "NagoBackend/utils"
    "net/http"

    loginForms "NagoBackend/forms/login"
    "github.com/labstack/echo/v4"
    //"github.com/k0kubun/pp"
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
    paramPasswordHash := utils.GetEncryptedHash(loginForm.Password)
    if paramPasswordHash != user.Password {
        return c.JSON(http.StatusOK, badRequestResponse([]string{"パスワードが一致しません"}))
    }
    authHandler := new(handlers.Auth)
    if err := authHandler.Login(c, user.ID); err != nil {
        return c.JSON(http.StatusOK, unauthorizedResponse([]string{"ログインに失敗しました"}))
    }
    return c.JSON(http.StatusOK, successResponse(map[string]interface{}{
        "user": user,
        "message": "ログインしました",
    }))
}

// ログアウト処理
func (lc *LoginController) Logout(c echo.Context) error {
    authHandler := new(handlers.Auth)
    if err := authHandler.Logout(c); err != nil {
        return c.JSON(http.StatusOK, serverErrorResponse([]string{"ログアウトに失敗しました"}))
    }
    return c.JSON(http.StatusOK, successResponse(map[string]string{
        "message": "ログアウトしました",
    }))
}
