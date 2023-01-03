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
			// 各種処理の前に設定する
			return next(cc)
		}
	})
}

func (cc *CustomContext) BindValidate(i interface{}) error {
	// bindとvalidateをまとめて処理する
	if err := cc.Bind(i); err != nil {
		cc.Logger().Error(err)
		return err
	}
	if err := cc.Validate(i); err != nil {
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
