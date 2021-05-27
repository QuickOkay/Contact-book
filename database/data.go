package database

import (
	"testcontact/models"

	"gorm.io/gorm"
)

func GetAllModels() []interface{} {
	return []interface{}{
		&models.Contact{},
	}
}

func MigrateDatabase(db *gorm.DB, migrate bool) {
	if db == nil {
		println("Database not connected")
	}

	if migrate {
		models := GetAllModels()
		db.AutoMigrate(models...)
	}
}
