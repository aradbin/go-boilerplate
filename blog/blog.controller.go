package blog

import (
	"app/database"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Create(c echo.Context) error {
	db := database.DB()
	item := new(Blog)

	if err := c.Bind(item); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request data")
	}

	result := db.Create(&item)
	if result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Something went wrong. Please try again")
	}

	return c.JSON(http.StatusCreated, item)
}

func GetAll(c echo.Context) error {
	db := database.DB()
	var items []Blog

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
	var item Blog

	hasItem, error := database.FindByField("id", c.Param("id"), &item)
	if error != nil {
		return error
	}

	return c.JSON(http.StatusOK, hasItem)
}

func Update(c echo.Context) error {
	db := database.DB()
	var item Blog
	updatedItem := new(Blog)

	if err := c.Bind(updatedItem); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request data")
	}

	hasItem, error := database.FindByField("id", c.Param("id"), &item)
	if error != nil && hasItem == nil {
		return error
	}

	result := db.Model(hasItem.(*Blog)).Updates(updatedItem)
	if result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Something went wrong. Please try again")
	}

	return c.JSON(http.StatusOK, hasItem)
}

func Delete(c echo.Context) error {
	db := database.DB()
	var item Blog

	hasItem, error := database.FindByField("id", c.Param("id"), &item)
	if error != nil && hasItem == nil {
		return error
	}

	result := db.Model(hasItem.(*Blog)).Delete(&Blog{})
	if result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Something went wrong. Please try again")
	}

	return c.JSON(http.StatusOK, result.RowsAffected)
}
