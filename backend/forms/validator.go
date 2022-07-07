package forms

import (
    "github.com/go-playground/validator"
    "github.com/labstack/echo/v4"
)

type Validator struct {
    validator *validator.Validate
}

func NewValidator() echo.Validator {
    return &Validator{validator: validator.New()}
}

func (v *Validator) Validate(i interface{}) error {
    return v.validator.Struct(i)
}

type Context struct {
    echo.Context
}

func (c *Context) BindValidate(i interface{}) error {
    if err := c.Bind(i); err != nil {
        return err
    }
    if err := c.Validate(i); err != nil {
        return err
    }
    return nil
}

//type callFunc func(c *Context) error
//
//func c(h callFunc) echo.HandlerFunc {
//    return func(c echo.Context) error {
//        return h(c.(*Context))
//    }
//}
