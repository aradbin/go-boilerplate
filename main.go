package main

import (
	"app/blog"
	"app/database"

	"github.com/labstack/echo/v4"
)

func main() {
	database.DatabaseInit()
	gorm := database.DB()
	gorm.AutoMigrate(&blog.Blog{})

	e := echo.New()
	blog.BlogRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
