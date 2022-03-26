package handler

import (
    "github.com/labstack/echo/v4"
)

func HandleGetPosts(c echo.Context) error {
    return c.String(200, "hello world")
}
