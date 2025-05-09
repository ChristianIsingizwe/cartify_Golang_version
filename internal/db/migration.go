package db

import (
	"carto/internal/utils"
	"database/sql"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
	"time"
)

var migrations = []*gormigrate.Migration{
	{
		ID: utils.MAKEID("create_users"),
		Migrate: func(tx *gorm.DB) error {
			if err := tx.Exec(`CREATE TYPE role AS ENUM('buyer', 'seller', 'admin')`).Error; err != nil {
				return err
			}

			type User struct {
				ID        uint `gorm:"primary_key"`
				CreatedAt time.Time
				UpdatedAt time.Time
				DeletedAt gorm.DeletedAt `gorm:"index"`

				FirstName    string `gorm:"size:255;not null"`
				LastName     string `gorm:"size:255;not null"`
				Email        string `gorm:"size:255;uniqueIndex;not null"`
				Role         string `gorm:"type:role;default:'buyer';not null"`
				PasswordHash string `gorm:"not null"`
				Age          uint8  `gorm:"not null"`
				ActivatedAt  sql.NullTime
			}
			if err := tx.AutoMigrate(&User{}).Error(); err != nil {
				return err
			}
			return nil
		},
	},
}
