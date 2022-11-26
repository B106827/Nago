package contexts

import (
    "errors"
    "fmt"
    "regexp"
    "strings"

    validation "github.com/go-ozzo/ozzo-validation/v4"
    "github.com/labstack/echo/v4"
)

type CustomContext struct {
    echo.Context
}

func InitCustomContext(e *echo.Echo) {
    e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            cc := &CustomContext{c}
            return next(cc)
        }
    })
}

func (c *CustomContext) BindValidate(i interface{}) error {
    if err := c.Bind(i); err != nil {
        c.Logger().Error(err)
        return err
    }
    if err := c.Validate(i); err != nil {
        errs := err.(validation.Errors)
        re := regexp.MustCompile(`{[a-z]+}`)
        for k, err := range errs {
            verr := err.(validation.Error)
            params := verr.Params()
            errs[k] = errors.New(re.ReplaceAllStringFunc(verr.Message(), func(s string) string {
                return fmt.Sprint(params[strings.Trim(s, "{}")])
            }))
        }
        return errs
    }
    return nil
}
