package handlers

import "github.com/labstack/echo/v4"

type UserHandler struct {
}

func (h *UserHandler) GetUser(c echo.Context) error {
	c.JSON(200, map[string]string{"message": "User retrieved successfully"})
	return nil
}
