package middlewares

import (
	"NagoBackend/config"
	"NagoBackend/constants"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitMiddleware(e *echo.Echo) {
	conf := config.GetConfig()
	// ログ
	e.Use(middleware.Logger())
	// エラーに対するリカバリー
	e.Use(middleware.Recover())
	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{conf.GetString("domain.front"), conf.GetString("domain.server")},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	}))
	// CSRF
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		ContextKey:     constants.CSRF_CONTEXT_KEY_NAME,
		CookieName:     constants.CSRF_COOKIE_KEY_NAME,
		CookieMaxAge:   constants.TIME_SECONDS_A_DAY,
		CookieSecure:   true,
		CookieHTTPOnly: true,
		TokenLookup:    "cookie:" + constants.CSRF_COOKIE_KEY_NAME,
	}))
	// バリデータ
	e.Validator = &CustomValidator{}
}

type CustomValidator struct{}

func (cv *CustomValidator) Validate(i interface{}) error {
	if c, ok := i.(validation.Validatable); ok {
		return c.Validate()
	}
	return nil
}
