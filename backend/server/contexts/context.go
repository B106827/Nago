package contexts

import (
    "github.com/labstack/echo/v4"
)

type CustomContext struct {
    echo.Context
}

func (c *CustomContext) BindValidate(i interface{}) error {
    if err := c.Bind(i); err != nil {
        return err
    }
    if err := c.Validate(i); err != nil {
        return err
    }
    return nil
}

func InitCustomContext(e *echo.Echo) {
    e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            cc := &CustomContext{c}
            return next(cc)
        }
    })
}

