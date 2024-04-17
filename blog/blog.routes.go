package blog

import (
	"github.com/labstack/echo/v4"
)

func BlogRoutes(e *echo.Echo) {
	api := e.Group("/blogs")

	api.GET("", GetAll)
	api.GET("/:id", GetOne)
	api.POST("", Create)
	api.PUT("/:id", Update)
	api.DELETE("/:id", Delete)
}
