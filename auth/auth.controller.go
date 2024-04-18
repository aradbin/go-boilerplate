package auth

import (
	"app/database"
	"app/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Register(c echo.Context) error {
	result := user.Create(c)

	return result
}

func Login(c echo.Context) error {
	var item user.User

	hasItem, error := database.FindByField("email", c.Param("id"), &item)
	if error != nil {
		return error
	}

	return c.JSON(http.StatusOK, hasItem)
}
