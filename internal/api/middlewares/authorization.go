package middlewares

import (
	"regexp"
	"strings"

	"github.com/labstack/echo/v4"
)



func (m *Middleware) Authorization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		auth := c.Request().Header.Get("Authorization")
		reg := regexp.MustCompile(`Bearer[\s]`)

		auth = reg.ReplaceAllString(auth, "")
		if strings.TrimSpace(auth) != "" {


			next(c)
		} else {
			return echo.NewHTTPError(401, "Unauthorized")
		}
		return nil
	}
}
