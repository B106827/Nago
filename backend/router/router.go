package router

import (
    "NagoBackend/api"
    "github.com/labstack/echo/v4"
)

func New() *echo.Echo {
    e := echo.New()

    // set main routes
    api.MainGroup(e)

    return e
}
