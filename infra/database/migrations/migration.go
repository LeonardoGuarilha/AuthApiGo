package migrations

import (
	"auth-api/domain/entities"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(&entities.User{})
}
