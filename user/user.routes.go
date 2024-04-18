package user

import (
	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Echo) {
	api := e.Group("/users")

	api.GET("", GetAll)
	api.GET("/:id", GetOne)
	api.POST("", Create)
	api.PUT("/:id", Update)
	api.DELETE("/:id", Delete)
}
