package main

import (
	"app/blog"
	"app/database"
	"app/user"

	"github.com/labstack/echo/v4"
)

func main() {
	database.DatabaseInit()
	gorm := database.DB()
	gorm.AutoMigrate(&blog.Blog{}, &user.User{})

	e := echo.New()
	blog.BlogRoutes(e)
	user.UserRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
