package controllers

import (
    "github.com/labstack/echo/v4"
    "net/http"
    userForms "NagoBackend/forms/user"
)

type UserController struct{}

func NewUserController() *UserController {
    return new(UserController)
}

// ログイン処理
func (uc *UserController) Login(c echo.Context) (err error) {
    loginForm := new(userForms.LoginForm)
    //if err = cc.BindValidate(loginForm); err != nil {
    //    return err
    //}
    if err := c.Bind(loginForm); err != nil {
        return err
    }
    if err := c.Validate(loginForm); err != nil {
        return err
    }
    return c.JSON(http.StatusOK, successResponse("OKです"))
}
