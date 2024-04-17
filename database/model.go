package database

import (
	"gorm.io/gorm"
)

type Model struct {
	gorm.Model
	CreatedBy *uint
	UpdatedBy *uint
	DeletedBy *uint
}
