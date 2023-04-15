package controllers

import (
	"NagoBackend/config"
	"NagoBackend/handlers"
	"NagoBackend/models"
	"NagoBackend/server/contexts"
	"NagoBackend/utils"
	"net/http"

	registerEmailForms "NagoBackend/forms/register_email"

	"github.com/labstack/echo/v4"
)

type RegisterEmailController struct{}

// メールアドレス仮登録処理
func (rec *RegisterEmailController) Create(c echo.Context) error {
	registerEmailForm := new(registerEmailForms.RegisterEmailForm)
	cc := c.(*contexts.CustomContext)
	if err := cc.BindValidate(registerEmailForm); err != nil {
		return c.JSON(http.StatusOK, badRequestResponse(err))
	}
	um := models.User{}
	user, err := um.FindByEmail(registerEmailForm.Email)
	if err != nil || user != nil {
		// エラー発生もしくはすでに登録済み
		return c.JSON(http.StatusOK, serverErrorResponse([]string{"すでに登録済みのメールアドレスです"}))
	}

	// メールアドレス仮登録
	rndID, err := utils.GetRandomText(10)
	if err != nil {
		return c.JSON(http.StatusOK, serverErrorResponse([]string{"エラーが発生しました"}))
	}
	utm := models.UserTemporary{}
	utm.ID = rndID
	utm.Email = registerEmailForm.Email
	// 有効期限は1日後
	utm.ExpiredAt = utils.GetAddDateTime(1, "d")
	if err := utm.Create(); err != nil {
		return c.JSON(http.StatusOK, serverErrorResponse([]string{"エラーが発生しました"}))
	}
	// 本登録用URLメール送信
	mailHandler := handlers.Mail{}
	conf := config.GetConfig()
	url := conf.GetString("domain.front")
	path := "/register/" + utm.ID
	vars := map[string]string{"Url": url + path}
	files := []string{"views/mail/register_email.tpl", "views/mail/base.tpl"}
	if err := mailHandler.Send(utm.Email, "【nago】新規登録用URL", vars, files...); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusOK, serverErrorResponse([]string{"登録用メールが送信できませんでした"}))
	}
	return c.JSON(http.StatusOK, successResponse(map[string]string{
		"message": "登録用URLを送信しました。メールを確認してください（有効期限1日）",
	}))
}
