package controllers

import (
    "net/http"
    "NagoBackend/models"
    "NagoBackend/server/contexts"
    "time"

    registerForms "NagoBackend/forms/register"
    "github.com/labstack/echo/v4"
)

type RegisterController struct{}

func NewRegisterController() *RegisterController {
    return &RegisterController{}
}

// 登録URL確認
func (rc *RegisterController) UrlCheck(c echo.Context) error {
    registerUrlForm := new(registerForms.RegisterUrlForm)
    cc := c.(*contexts.CustomContext)
    if err := cc.BindValidate(registerUrlForm); err != nil {
        return c.JSON(http.StatusOK, badRequestResponse(err))
    }
    ut := new(models.UserTemporary)
    if err := ut.FindById(registerUrlForm.TmpID); err != nil {
        return c.JSON(http.StatusOK, badRequestResponse([]string{"無効なURLです"}))
    }
    if ut.ExpiredAt.Unix() < time.Now().Unix() {
        // 有効期限が切れていれば無効
        return c.JSON(http.StatusOK, badRequestResponse([]string{"有効期限(24時間)が切れています。もう一度登録をお願いします"}))
    }
    return c.JSON(http.StatusOK, successResponse(map[string]string{
        "email": ut.Email,
    }))
}

// 登録処理
func (rc *RegisterController) Register(c echo.Context) error {
    registerForm := new(registerForms.RegisterForm)
    cc := c.(*contexts.CustomContext)
    if err := cc.BindValidate(registerForm); err != nil {
        return err
    }
    return c.JSON(http.StatusOK, successResponse("OKです"))
}
