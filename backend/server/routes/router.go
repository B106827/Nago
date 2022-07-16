package routes

import (
    "github.com/labstack/echo/v4"
    "NagoBackend/controllers"
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
    api.POST("/register_url_check", registerController.URLCheck)
    // register
    api.POST("/register", registerController.Register)

    // login
    loginController := controllers.NewLoginController()
    api.POST("/login", loginController.Login)
    // logout
    api.GET("/logout", loginController.Logout)
}
