package models

import (
	"github.com/google/uuid"
)


// type Gender int

// const (
// 	MALE Gender = iota
// 	FEMALE
// )

type User struct {
	Id uuid.UUID `json:"id" gorm:"type:uuid;primary_key"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	Gender string `json:"gender"`
	Blogs []Blog `gorm:"foreignKey:UserId"`
}