package controllers

import (
    "github.com/labstack/echo/v4"
    "net/http"
    userForms "NagoBackend/forms/user"
    "NagoBackend/server/contexts"
)

type UserController struct{}

func NewUserController() *UserController {
    return new(UserController)
}

// ログイン処理
func (uc *UserController) Login(c echo.Context) (err error) {
    loginForm := new(userForms.LoginForm)
    cc := c.(*contexts.CustomContext)
    if err := cc.BindValidate(loginForm); err != nil {
        return err
    }
    return c.JSON(http.StatusOK, successResponse("OKです"))
}
