package controllers

import (
	"NagoBackend/handlers"
	"NagoBackend/models"
	"NagoBackend/server/contexts"
	"NagoBackend/utils"
	"net/http"

	loginForms "NagoBackend/forms/login"

	"github.com/labstack/echo/v4"
)

type LoginController struct{}

// ログイン処理
func (lc *LoginController) Create(c echo.Context) error {
	loginForm := new(loginForms.LoginForm)
	cc := c.(*contexts.CustomContext)
	if err := cc.BindValidate(loginForm); err != nil {
		return c.JSON(http.StatusOK, badRequestResponse(err))
	}
	um := models.User{}
	user, err := um.FindByEmail(loginForm.Email)
	if err != nil || user == nil {
		// エラーもしくはユーザーが存在しない
		return c.JSON(http.StatusOK, notFoundResponse([]string{"メンバーが見つかりません"}))
	}
	paramPasswordHash := utils.GetEncryptedHash(loginForm.Password)
	if paramPasswordHash != user.Password {
		return c.JSON(http.StatusOK, badRequestResponse([]string{"パスワードが一致しません"}))
	}
	authHandler := handlers.Auth{}
	if err := authHandler.Login(c, user.ID); err != nil {
		return c.JSON(http.StatusOK, unauthorizedResponse([]string{"ログインに失敗しました"}))
	}
	cm := models.Cart{}
	cartList, err := cm.FindByUserIdWithProduct(user.ID)
	if err != nil {
		c.Logger().Error(err)
	}
	return c.JSON(http.StatusOK, successResponse(map[string]interface{}{
		"user":     user,
		"cartList": cartList,
		"message":  "ログインしました",
	}))
}

// ログアウト処理
func (lc *LoginController) Logout(c echo.Context) error {
	authHandler := handlers.Auth{}
	if err := authHandler.Logout(c); err != nil {
		return c.JSON(http.StatusOK, serverErrorResponse([]string{"ログアウトに失敗しました"}))
	}

	return c.JSON(http.StatusOK, successResponse(map[string]string{
		"message": "ログアウトしました",
	}))
}
