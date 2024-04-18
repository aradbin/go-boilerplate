package user

import (
	"app/database"
)

type User struct {
	database.Model
	Username string
	Email    string
	Password string
}
