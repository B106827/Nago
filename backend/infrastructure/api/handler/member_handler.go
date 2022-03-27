package handler

import (
    "context"
    "github.com/labstack/echo/v4"
    "NagoBackend/domain/model"
    "NagoBackend/interface/controllers"
    "net/http"
)

type memberHandler struct {
    memberController controllers.MemberController
}

type MemberHandler interface {
    GetMembers(c echo.Context) error
}

func NewMemberHandler(uc controllers.MemberController) MemberHandler {
    return &memberHandler{memberController: uc}
}

func (uh *memberHandler) GetMembers(c echo.Context) error {
    req := &model.Member{}
    if err := c.Bind(req); err != nil {
        return c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
    }

    ctx := c.Request().Context()
    if ctx == nil {
        ctx = context.Background()
    }

    u, err := uh.memberController.GetMembers()
    if err != nil {
        // システム内のエラー
        return c.JSON(http.StatusBadRequest, err)
    }
    return c.JSON(http.StatusOK, u)
}
