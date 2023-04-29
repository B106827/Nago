package controllers

import (
	"NagoBackend/constants"
	"net/http"

	"NagoBackend/models"
	"NagoBackend/server/contexts"
	"NagoBackend/utils"
	"time"

	registerForms "NagoBackend/forms/register"

	"github.com/labstack/echo/v4"
)

type RegisterController struct{}

// 登録URL確認
func (rc *RegisterController) UrlCheck(c echo.Context) error {
	registerUrlForm := new(registerForms.RegisterUrlForm)
	cc := c.(*contexts.CustomContext)
	if err := cc.BindValidate(registerUrlForm); err != nil {
		return c.JSON(http.StatusOK, badRequestResponse(err))
	}
	utm := models.UserTemporary{}
	ut, err := utm.FindById(registerUrlForm.TmpID)
	if err != nil || ut == nil {
		// エラーもしくはデータが存在しない
		return c.JSON(http.StatusOK, notFoundResponse([]string{"無効なURLです"}))
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
func (rc *RegisterController) Create(c echo.Context) error {
	registerForm := new(registerForms.RegisterForm)
	cc := c.(*contexts.CustomContext)
	if err := cc.BindValidate(registerForm); err != nil {
		return c.JSON(http.StatusOK, badRequestResponse(err))
	}
	utm := models.UserTemporary{}
	ut, err := utm.FindById(registerForm.TmpID)
	if err != nil {
		return c.JSON(http.StatusOK, notFoundResponse([]string{"無効なURLです"}))
	}
	if ut.ExpiredAt.Unix() < time.Now().Unix() {
		// 有効期限が切れていれば無効
		return c.JSON(http.StatusOK, badRequestResponse([]string{"有効期限(24時間)が切れています。もう一度登録をお願いします"}))
	}
	um := models.User{}
	user, err := um.FindByEmail(registerForm.Email)
	if err != nil || user != nil {
		return c.JSON(http.StatusOK, badRequestResponse([]string{"すでに登録済みのメールアドレスです"}))
	}
	um.Email = registerForm.Email
	um.Password = utils.GetEncryptedHash(registerForm.Password)
	um.Name = registerForm.Name
	um.Status = constants.USER_STATUS_REGISTERED
	if err := um.Create(); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusOK, serverErrorResponse([]string{"エラーが発生しました"}))
	}
	// 登録が完了すれば一時テーブルから削除しておく
	if err := ut.Delete(); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusOK, serverErrorResponse([]string{"エラーが発生しました"}))
	}
	return c.JSON(http.StatusOK, successResponse("OKです"))
}
