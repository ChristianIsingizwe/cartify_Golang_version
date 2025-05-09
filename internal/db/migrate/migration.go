package migrate

import (
	"carto/internal/enums"
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

			if err := tx.Exec(`CREATE TYPE IF NOT EXISTS role AS ENUM('buyer','seller','admin');`).Error; err != nil {
				return err
			}

			type User struct {
				ID        uint `gorm:"primaryKey"`
				CreatedAt time.Time
				UpdatedAt time.Time
				DeletedAt gorm.DeletedAt `gorm:"index"`

				FirstName    string
				LastName     string
				Email        string `gorm:"unique"`
				PasswordHash string
				Role         enums.UserRoles `gorm:"type:role;default"`
				ActivatedAt  sql.NullTime
			}
			if err := tx.AutoMigrate(&User{}); err != nil {
				return err
			}
			return nil
		},

		Rollback: func(tx *gorm.DB) error {
			if err := tx.Migrator().DropTable("users"); err != nil {
				return err
			}

			if err := tx.Exec("DROP TYPE IF EXISTS role CASCADE").Error; err != nil {
				return err
			}
			return nil
		},
	},
}
