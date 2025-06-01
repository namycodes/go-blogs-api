package blog

import "gorm.io/gorm"

type Blog struct {
	gorm.Model
	Title string `json:"title" gorm:"unique"`
	Description string `json:"description"`
}