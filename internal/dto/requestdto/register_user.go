package dto

type RegisterUserDto struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName string `json:"last_name" binding:"required"`
	Email string `json:"email" binding:"required"`
	Gender string `json:"gender" binding:"required,oneof=MALE FEMALE"`
	Password string `json:"password" binding:"required"`
}