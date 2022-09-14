package entity

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID         uint    `gorm:"unique,primaryKey,id"`
	Name       string  `gorm:""`
	Point      float32 `gorm:""`
	VotesCount uint64  `gorm:""`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
