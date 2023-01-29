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
	authMiddleware := middlewares.AuthMiddleware{}

	/*
	  /api
	*/
	api := e.Group("/api")
	//
	// ヘルスチェック
	//
	healthController := controllers.HealthController{}
	api.GET("/health", healthController.Index)
	//
	// 新規登録のためのメール送信
	//
	registerEmailController := controllers.RegisterEmailController{}
	api.POST("/register_email", registerEmailController.RegisterEmail)
	//
	// 新規登録
	//
	registerController := controllers.RegisterController{}
	api.POST("/register_url_check", registerController.UrlCheck)
	api.POST("/register", registerController.Register)
	//
	// ログイン・ログアウト
	//
	loginController := controllers.LoginController{}
	api.POST("/login", loginController.Login)
	api.GET("/logout", loginController.Logout, authMiddleware.Authenticate())

	/*
	  /api/product
	*/
	apiP := e.Group("/api/product")
	//
	// 商品一覧
	//
	productController := controllers.ProductController{}
	apiP.GET("", productController.Index)
	//
	// 商品情報
	//
	apiP.GET("/:productId", productController.Get)

	/*
	  /api/user
	*/
	apiU := e.Group("/api/user", authMiddleware.Authenticate())
	//
	// ユーザー情報
	//
	userMyinfoController := controllers.UserMyinfoController{}
	apiU.GET("/myinfo", userMyinfoController.Index)
}
