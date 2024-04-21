package user

import (
	"app/database"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("your_secret_key")

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

	hashedPassword, error := HashPassword(newItem.Password)
	if error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Something went wrong. Please try again")
	}
	newItem.Password = string(hashedPassword)

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

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func CheckPasswordHash(password string, hashedPassword []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	return err == nil
}

func GenerateJWT(user User) (string, error) {
	expirationTime := time.Now().Add(time.Hour * 1000) // Set token expiration time (e.g., 1 hour)
	claims := jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"exp":      expirationTime.Unix(), // Expiration time in seconds
		"issuedAt": time.Now().Unix(),     // Issued at time in seconds
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
