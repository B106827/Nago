package server

import(
    "github.com/go-playground/validator"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "NagoBackend/config"
    "NagoBackend/controllers"
    "net/http"
)

type (
    CustomValidator struct {
        validator *validator.Validate
    }
)

func (cv *CustomValidator) Validate(i interface{}) error {
    if err := cv.validator.Struct(i); err != nil {
        return err
    }
    return nil
}

func NewRouter() (*echo.Echo, error) {
    c := config.GetConfig()
    router := echo.New()
    router.Use(middleware.Logger())
    router.Use(middleware.Recover())
    router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
        AllowOrigins: c.GetStringSlice("server.cors"),
        AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
    }))
    router.Validator = &CustomValidator{validator: validator.New()}

    version := router.Group("/" + c.GetString("server.version"))

    // health
    healthController := controllers.NewHealthController()
    version.GET("/health", healthController.Index)

    // user
    userController := controllers.NewUserController()
    version.POST("/login", userController.Login)

    return router, nil
}
