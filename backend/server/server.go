package server

import (
    "github.com/labstack/echo/v4"
    "NagoBackend/config"
    "NagoBackend/server/routes"
    "NagoBackend/server/middlewares"
    "NagoBackend/server/contexts"
)

func Init() error {
    e := echo.New()
    // initialize context
    contexts.InitCustomContext(e)

    // initialize router
    routes.InitRouter(e)
    // initialize middleware
    middlewares.InitMiddleware(e)

    c := config.GetConfig()
    // server start
    e.Logger.Fatal(e.Start(":" + c.GetString("server.port")))
    return nil
}
