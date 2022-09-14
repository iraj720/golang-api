package fixture

import (
	"snapp/internal/entity"

	"gorm.io/gorm"
)

func ProductTable(db *gorm.DB) {
	products := []entity.Product{
		{
			ID:   1,
			Name: "TV",
		},
		{
			ID:   2,
			Name: "carpet",
		},
		{
			ID:   3,
			Name: "car",
		},
		{
			ID:   4,
			Name: "pc",
		},
	}
	for _, val := range products {
		db.Model(&entity.Product{}).Create(&val)
	}
}
