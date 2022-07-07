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

    // user
    userController := controllers.NewUserController()
    version.POST("/login", userController.Login)
}
