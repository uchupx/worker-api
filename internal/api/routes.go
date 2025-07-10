package api

import (
	"github.com/labstack/echo/v4"
	"github.com/uchupx/worker-api/internal"
	"github.com/uchupx/worker-api/internal/config"
)

func InitRouter(conf *config.Config, e *echo.Echo) {
	factory := internal.Factory{}


	factory.GetDBConnection()

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"message": "Healthy"})
	})

	userHandler := factory.GetUserHandler()

	e.GET("/user/", userHandler.GetUser)
}
