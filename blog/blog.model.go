package blog

import (
	"app/database"
)

type Blog struct {
	database.Model
	Title       string
	Description string
}
