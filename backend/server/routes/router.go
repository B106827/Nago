package routes

import (
	"github.com/labstack/echo/v4"
	"NagoBackend/config"
	"NagoBackend/constants"
	"NagoBackend/controllers"
	"NagoBackend/handlers"

	"github.com/labstack/echo/v4/middleware"
)

func InitRouter(e *echo.Echo) {
	api := e.Group("/api")
	// health
	healthController := controllers.NewHealthController()
	api.GET("/health", healthController.Index)

	// register email
	registerEmailController := controllers.NewRegisterEmailController()
	api.POST("/register_email", registerEmailController.RegisterEmail)

	// register url check
	registerController := controllers.NewRegisterController()
	api.POST("/register_url_check", registerController.UrlCheck)
	// register
	api.POST("/register", registerController.Register)

	// login
	loginController := controllers.NewLoginController()
	api.POST("/login", loginController.Login)
	// logout
	api.GET("/logout", loginController.Logout)

	// cookie で管理されている JWT をチェックする
	conf := config.GetConfig()
	jwtConfig := middleware.JWTConfig{
		Claims:     &handlers.JwtCustomClaims{},
		SigningKey: []byte(conf.GetString("session.secret")),
		TokenLookup: "cookie:" + constants.SESSION_COOKIE_KEY_NAME,
	}
	// user
	apiU := e.Group("/api/user")
	apiU.Use(middleware.JWTWithConfig(jwtConfig))
	// mypage
	userMyinfoController := controllers.NewUserMyinfoController()
	apiU.GET("/myinfo", userMyinfoController.Index)
}
