package migrations

import (
	"snapp/internal/entity"

	"gorm.io/gorm"
)

func MigrateUp(db *gorm.DB) {
	db.AutoMigrate(&entity.Product{})
}

func MigrateDown(db *gorm.DB) {

}
