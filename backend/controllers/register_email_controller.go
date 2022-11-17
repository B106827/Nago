package controllers

import (
    "NagoBackend/config"
    "NagoBackend/models"
    "NagoBackend/server/contexts"
    "NagoBackend/services"
    "NagoBackend/utils"
    "net/http"
    "time"

    registerEmailForms "NagoBackend/forms/register_email"
    "github.com/labstack/echo/v4"
)

type RegisterEmailController struct{}

func NewRegisterEmailController() *RegisterEmailController {
    return &RegisterEmailController{}
}

// メールアドレス仮登録処理
func (rec *RegisterEmailController) RegisterEmail(c echo.Context) error {
    registerEmailForm := new(registerEmailForms.RegisterEmailForm)
    cc := c.(*contexts.CustomContext)
    if err := cc.BindValidate(registerEmailForm); err != nil {
        return c.JSON(http.StatusOK, badRequestResponse(err))
    }
    user := new(models.User)
    if err := user.FindByEmail(registerEmailForm.Email); err == nil {
        return c.JSON(http.StatusOK, badRequestResponse([]string{"すでに登録済みのメールアドレスです"}))
    }

    // メールアドレス仮登録
    rndID, err := utils.GetRandomText(10)
    if err != nil {
        c.Logger().Error(err)
        return c.JSON(http.StatusOK, serverErrorResponse([]string{"エラーが発生しました"}))
    }
    ut := new(models.UserTemporary)
    ut.ID    = rndID
    ut.Email = registerEmailForm.Email
    // 有効期限は1日後
    ut.ExpiredAt = utils.GetAddDateTime(1, "d")
    ut.CreatedAt = time.Now()
    if err := ut.Create(); err != nil {
        c.Logger().Error(err)
        return c.JSON(http.StatusOK, serverErrorResponse([]string{"エラーが発生しました"}))
    } 
    // 本登録用URLメール送信
    mail  := services.NewMail()
    conf  := config.GetConfig()
    url   := conf.GetString("url.main")
    path  := "/register/" + ut.ID
    vars  := map[string]string{"Url": url + path}
    files := []string{"views/mail/register_email.tpl", "views/mail/base.tpl"}
    if err := mail.Send(ut.Email, "【nago】新規登録用URL", vars, files...); err != nil {
        c.Logger().Error(err)
        return c.JSON(http.StatusOK, serverErrorResponse([]string{"登録用メールが送信できませんでした"}))
    }
    return c.JSON(http.StatusOK, successResponse(map[string]string{
        "message": "登録用URLを送信しました。メールを確認してください（有効期限1日）",
    }))
}
