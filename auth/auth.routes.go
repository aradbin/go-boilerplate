package auth

import (
	"github.com/labstack/echo/v4"
)

func AuthRoutes(e *echo.Echo) {
	api := e.Group("/auth")

	api.POST("/register", Register)
	api.POST("/login", Login)
}
