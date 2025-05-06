package database

import (
	"gorm-in-go/pkg/models"
	"log"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {

	// Auto Migrate
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Error migrating database: ", err)
	}

	return nil
}
