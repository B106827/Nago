package routes

import (
	"NagoBackend/controllers"
	"NagoBackend/server/middlewares"

	"github.com/labstack/echo/v4"
)

func InitRouter(e *echo.Echo) {
	/*
	  認証ミドルウェア
	*/
	authMiddleware := middlewares.NewAuthMiddleware()

	/*
	  /api
	*/
	api := e.Group("/api")
	//
	// ヘルスチェック
	//
	healthController := controllers.NewHealthController()
	api.GET("/health", healthController.Index)
	//
	// 新規登録のためのメール送信
	//
	registerEmailController := controllers.NewRegisterEmailController()
	api.POST("/register_email", registerEmailController.RegisterEmail)
	//
	// 新規登録
	//
	registerController := controllers.NewRegisterController()
	api.POST("/register_url_check", registerController.UrlCheck)
	api.POST("/register", registerController.Register)
	//
	// ログイン・ログアウト
	//
	loginController := controllers.NewLoginController()
	api.POST("/login", loginController.Login)
	api.GET("/logout", loginController.Logout, authMiddleware.Authenticate())
	/*
	  /api/user
	*/
	apiU := e.Group("/api/user", authMiddleware.Authenticate())
	//
	// ユーザー情報
	//
	userMyinfoController := controllers.NewUserMyinfoController()
	apiU.GET("/myinfo", userMyinfoController.Index)
}
