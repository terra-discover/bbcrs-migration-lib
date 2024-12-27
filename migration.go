package migration

import (
	"gorm.io/gorm"
)

// AutoMigrate all tables
func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(ModelMigrations...)
}
