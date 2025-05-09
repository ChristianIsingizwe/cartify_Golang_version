package models

import (
	"carto/internal/enums"
	"database/sql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	FirstName    string
	LastName     string
	Email        string          `gorm:"unique"`
	Role         enums.UserRoles `gorm:"type:role;default:'buyer'"`
	PasswordHash string
	Age          uint8
	ActivatedAt  sql.NullTime
}
