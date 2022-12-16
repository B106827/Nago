package controllers

import (
    "net/http"

//    "github.com/golang-jwt/jwt"
    "NagoBackend/handlers"
//    "NagoBackend/server/contexts"
//    "NagoBackend/utils"
//    "net/http"
//
//    loginForms "NagoBackend/forms/login"
    "github.com/labstack/echo/v4"
//    "github.com/k0kubun/pp"
)

type UserMyinfoController struct{}

func NewUserMyinfoController() *UserMyinfoController {
    return &UserMyinfoController{}
}

func (umc *UserMyinfoController) Index(c echo.Context) error {
    //user := c.Get("user").(*jwt.Token)
    // TODO: ミドルウェアでユーザー情報を取得してどこからでもとれるように
    //claims := user.Claims.(*handlers.JwtCustomClaims)
    //pp.Print(claims)
	//userID := claims.UserID
	authHandler := new(handlers.Auth)
	authHandler.Test(c)
    return c.JSON(http.StatusOK, successResponse(map[string]interface{}{
        "message": "success",
    }))
}
