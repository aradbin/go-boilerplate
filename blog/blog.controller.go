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
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	result := db.Create(&item)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": result.Error.Error(),
		})
	}

	return c.JSON(http.StatusCreated, item)
}

func GetAll(c echo.Context) error {
	db := database.DB()
	var items []Blog

	query, total := database.Query(c, db, &items)
	if err := query.Find(&items).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"results": items,
		"total":   total,
	})
}

func GetOne(c echo.Context) error {
	db := database.DB()
	id := c.Param("id")
	var item Blog

	result := db.Where("id = ?", id).First(&item)
	if result.Error != nil {
		if result.RowsAffected == 0 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "Item not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, item)
}

func Update(c echo.Context) error {
	db := database.DB()
	id := c.Param("id")
	var item Blog

	result := db.Where("id = ?", id).First(&item)
	if result.Error != nil {
		if result.RowsAffected == 0 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "Item not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": result.Error.Error(),
		})
	}

	updatedItem := new(Blog)
	if err := c.Bind(updatedItem); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	result = db.Model(&item).Updates(updatedItem)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, item)
}

func Delete(c echo.Context) error {
	db := database.DB()
	id := c.Param("id")

	result := db.Where("id = ?", id).Delete(&Blog{})
	if result.Error != nil {
		if result.RowsAffected == 0 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "Item not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": result.Error.Error(),
		})
	}

	return c.JSON(http.StatusNoContent, nil)
}
