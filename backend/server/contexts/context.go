package contexts

import (
    "errors"
    "fmt"
    "regexp"
    "strings"

    validation "github.com/go-ozzo/ozzo-validation/v4"
    "github.com/labstack/echo/v4"

//    "github.com/k0kubun/pp"
)

type CustomContext struct {
    echo.Context
}

type CustomError struct {
    errorType     string
    originalError error
}

func (c *CustomContext) BindValidate(i interface{}) error {
    if err := c.Bind(i); err != nil {
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

func InitCustomContext(e *echo.Echo) {
    e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            cc := &CustomContext{c}
            return next(cc)
        }
    })
}

