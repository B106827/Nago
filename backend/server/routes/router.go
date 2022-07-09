package routes

import (
    "github.com/labstack/echo/v4"
    "NagoBackend/config"
    "NagoBackend/controllers"
)

func InitRouter(e *echo.Echo) {
    c := config.GetConfig()
    version := e.Group("/" + c.GetString("server.version"))

    // health
    healthController := controllers.NewHealthController()
    version.GET("/health", healthController.Index)

    // register email
    registerEmailController := controllers.NewRegisterEmailController()
    version.POST("/register_email", registerEmailController.RegisterEmail)

    // register url check
    registerController := controllers.NewRegisterController()
    version.POST("/register_url_check", registerController.URLCheck)
    // register
    version.POST("/register", registerController.Register)

    // login
    loginController := controllers.NewLoginController()
    version.POST("/login", loginController.Login)
    // logout
    version.GET("/logout", loginController.Logout)
}
