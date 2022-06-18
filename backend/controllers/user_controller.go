package controllers

import (
    "github.com/labstack/echo/v4"
    "net/http"
    userForms "NagoBackend/forms/user"
)

// UserController is controller for login
type UserController struct{}

func NewUserController() *UserController {
    return new(UserController)
}

// User is handle login process
func (uc *UserController) Login(c echo.Context) (err error) {
    loginForm := new(userForms.LoginForm)
    if err = c.Bind(loginForm); err != nil {
        return c.JSON(http.StatusOK, newResponse(
            http.StatusBadRequest,
            http.StatusText(http.StatusBadRequest),
            "bind error",
        ))
    }
    if err = c.Validate(loginForm); err != nil {
        return c.JSON(http.StatusOK, newResponse(
            http.StatusBadRequest,
            http.StatusText(http.StatusBadRequest),
            "validate error",
        ))
    }
    return c.JSON(http.StatusOK, newResponse(
        http.StatusOK,
        http.StatusText(http.StatusOK),
        "validation ok",
    ))
}
