package server

import(
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "NagoBackend/config"
    "NagoBackend/controllers"
    "net/http"
)

func NewRouter() (*echo.Echo, error) {
    c := config.GetConfig()
    router := echo.New()
    router.Use(middleware.Logger())
    router.Use(middleware.Recover())
    router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
        AllowOrigins: c.GetStringSlice("server.cors"),
        AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
    }))

    version := router.Group("/" + c.GetString("server.version"))

    healthController := controllers.NewHealthController()
    version.GET("/health", healthController.Index)

    return router, nil
}
