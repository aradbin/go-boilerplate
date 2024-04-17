package database

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB
var e error

func DatabaseInit() {
	host := "127.0.0.1"
	user := "db"
	password := "db"
	dbName := "db"
	port := 5432

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", host, user, password, dbName, port)
	database, e = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if e != nil {
		panic(e)
	}
}

func DB() *gorm.DB {
	return database
}

func Query(c echo.Context, db *gorm.DB, model interface{}) (*gorm.DB, int64) {
	db = Sort(c, db).Model(model)

	var total int64
	db.Count(&total)

	db = Paginate(c, db)

	return db, total
}

func Sort(c echo.Context, db *gorm.DB) *gorm.DB {
	sort := "created_at"
	order := "desc"
	if c.QueryParam("sort") != "" {
		sort = c.QueryParam("sort")
	}
	if c.QueryParam("order") != "" {
		order = c.QueryParam("order")
	}
	db = db.Order(sort + " " + order)

	return db
}

func Paginate(c echo.Context, db *gorm.DB) *gorm.DB {
	if c.QueryParam("limit") != "" {
		limit, err := strconv.Atoi(c.QueryParam("limit"))
		if err == nil && limit > 0 {
			db = db.Limit(limit)
			if c.QueryParam("page") != "" {
				page, err := strconv.Atoi(c.QueryParam("page"))
				if err == nil && page > 0 {
					offset := (page - 1) * limit
					db = db.Offset(offset)
				}
			}
		}
	}

	return db
}
