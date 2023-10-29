package server

import (
	"NagoBackend/config"
	"NagoBackend/server/contexts"
	"NagoBackend/server/middlewares"
	"NagoBackend/server/routes"
	"fmt"

	"github.com/labstack/echo/v4"
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
	e.Logger.Error(e.Start(fmt.Sprintf(":%d", c.GetInt("server.port"))))
	return nil
}
