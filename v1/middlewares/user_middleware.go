package middlewares

import "github.com/labstack/echo/v4"

type (
	UserMiddleware struct{}
)

func (m UserMiddleware) ValidateToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return next(c)
	}
}
