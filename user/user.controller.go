package user

import (
	"app/database"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Create(c echo.Context) error {
	db := database.DB()
	var item User
	newItem := new(User)

	if err := c.Bind(newItem); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request data")
	}

	hasItemByEmail, error := database.FindByField("email", newItem.Email, &item)
	if error == nil && hasItemByEmail != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Email already exists")
	}

	hasItemByUsername, error := database.FindByField("username", newItem.Username, &item)
	if error == nil && hasItemByUsername != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Username already exists")
	}

	result := db.Create(&newItem)
	if result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Something went wrong. Please try again")
	}

	return c.JSON(http.StatusCreated, newItem)
}

func GetAll(c echo.Context) error {
	db := database.DB()
	var items []User

	query, total := database.Query(c, db, &items)
	result := query.Find(&items)
	if result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Something went wrong. Please try again")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"results": items,
		"total":   total,
	})
}

func GetOne(c echo.Context) error {
	var item User

	hasItem, error := database.FindByField("id", c.Param("id"), &item)
	if error != nil {
		return error
	}

	return c.JSON(http.StatusOK, hasItem)
}

func Update(c echo.Context) error {
	db := database.DB()
	var item User

	updatedItem := new(User)
	if err := c.Bind(updatedItem); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request data")
	}

	hasItem, error := database.FindByField("id", c.Param("id"), &item)
	if error != nil && hasItem == nil {
		return error
	}

	result := db.Model(hasItem.(*User)).Updates(updatedItem)
	if result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Something went wrong. Please try again")
	}

	return c.JSON(http.StatusOK, hasItem)
}

func Delete(c echo.Context) error {
	db := database.DB()
	var item User

	hasItem, error := database.FindByField("id", c.Param("id"), &item)
	if error != nil && hasItem == nil {
		return error
	}

	result := db.Model(hasItem.(*User)).Delete(&User{})
	if result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Something went wrong. Please try again")
	}

	return c.JSON(http.StatusOK, result.RowsAffected)
}
