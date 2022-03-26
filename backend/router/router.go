package router

import (
    "NagoBackend/handlers"
    "github.com/labstack/echo/v4"
)

func New() *echo.Echo {
    e := echo.New()

    // set main routes
    handlers.MainGroup(e)

    return e
}
