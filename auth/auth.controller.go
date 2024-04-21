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
	newItem := new(user.User)
	var item user.User
	var hasItem interface{}
	var error error

	if err := c.Bind(newItem); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request data")
	}

	if newItem.Username != "" {
		hasItem, error = database.FindByField("username", newItem.Username, &item)
		if error != nil {
			return echo.NewHTTPError(http.StatusNotFound, "User not registered. Please register")
		}
	} else {
		hasItem, error = database.FindByField("email", newItem.Email, &item)
		if error != nil {
			return echo.NewHTTPError(http.StatusNotFound, "User not registered. Please register")
		}
	}

	hasUser := hasItem.(*user.User)

	if !user.CheckPasswordHash(newItem.Password, []byte(hasUser.Password)) {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid username or password")
	}

	token, err := user.GenerateJWT(*hasUser)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Something went wrong. Please try again")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
		"user":  hasUser,
	})
}
