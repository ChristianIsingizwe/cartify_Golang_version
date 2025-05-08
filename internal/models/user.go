package models

import (
	"database/sql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	FirstName   string
	LastName    string
	Email       string
	Age         uint8
	Location    string
	ActivatedAt sql.NullTime
}
