package models

import (
	"github.com/google/uuid"
)

type Blog struct {
    Id uuid.UUID `gorm:"type:uuid;primaryKey"`
	Title string `json:"title" gorm:"unique"`
	Description string `json:"description"`
	UserId uuid.UUID `gorm:"type:uuid;index"`
	User User `gorm:"foreignKey:UserId"`
}