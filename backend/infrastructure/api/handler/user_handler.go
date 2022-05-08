package handler

import (
    "context"
    "github.com/labstack/echo/v4"
    "NagoBackend/domain/model"
    "NagoBackend/interface/controllers"
    "net/http"
)

type userHandler struct {
    userController controllers.UserController
}

type UserHandler interface {
    GetUsers(c echo.Context) error
}

func NewUserHandler(uc controllers.UserController) UserHandler {
    return &userHandler{userController: uc}
}

func (uh *userHandler) GetUsers(c echo.Context) error {
    req := &model.User{}
    if err := c.Bind(req); err != nil {
        return c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
    }

    ctx := c.Request().Context()
    if ctx == nil {
        ctx = context.Background()
    }

    u, err := uh.userController.GetUsers()
    if err != nil {
        // システム内のエラー
        return c.JSON(http.StatusBadRequest, err)
    }
    return c.JSON(http.StatusOK, u)
}
