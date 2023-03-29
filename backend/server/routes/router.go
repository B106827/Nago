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

	// ヘルスチェック
	healthController := controllers.HealthController{}
	api.GET("/health", healthController.Index)

	// 新規登録のためのメール送信
	registerEmailController := controllers.RegisterEmailController{}
	api.POST("/register_email", registerEmailController.Create)

	// 新規登録
	registerController := controllers.RegisterController{}
	api.POST("/register_url_check", registerController.UrlCheck)
	api.POST("/register", registerController.Create)

	// ログイン・ログアウト
	loginController := controllers.LoginController{}
	api.POST("/login", loginController.Create)
	api.GET("/logout", loginController.Logout, authMiddleware.Authenticate())

	/*
	  /api/product
	*/
	apiP := e.Group("/api/product")

	// 商品一覧
	productController := controllers.ProductController{}
	apiP.GET("", productController.Index)
	// 商品情報
	apiP.GET("/:productId", productController.Get)

	/*
	  /api/user
	*/
	apiU := e.Group("/api/user", authMiddleware.Authenticate())

	// ユーザー情報
	userMyinfoController := controllers.UserMyinfoController{}
	apiU.GET("/myinfo", userMyinfoController.Index)

	/*
	  /api/cart
	*/
	apiC := e.Group("/api/cart", authMiddleware.Authenticate())

	// カート更新
	cartController := controllers.CartController{}
	apiC.PUT("", cartController.Update)
	// カート削除
	apiC.DELETE("", cartController.Delete)

	/*
	  /api/order
	*/
	apiO := e.Group("/api/order", authMiddleware.Authenticate())

	// 決済処理
	orderController := controllers.OrderController{}
	apiO.POST("/create", orderController.Create)

	/*
	  /api/master
	*/
	apiM := e.Group("/api/master")

	// 都道府県一覧
	masterController := controllers.MasterController{}
	apiM.GET("/pref", masterController.GetPrefMaster)
}
