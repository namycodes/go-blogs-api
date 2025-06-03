package authentication

import (
	dto "com.namycodes/internal/dto/requestdto"
	"com.namycodes/internal/models"
	"gorm.io/gorm"
)


type Repository interface{
	Create(user *models.User) (*models.User, error)
	Login(user *dto.AuthenticationDto) (*models.User, error)
}

type repositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository{
	return &repositoryImpl{db: db}
}

func (r *repositoryImpl) Create(user *models.User) (*models.User, error){
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *repositoryImpl) Login(user *dto.AuthenticationDto) (*models.User, error){
	var dbUser *models.User
	if err := r.db.Where("email =? ", user.Email).First(&dbUser).Error; err != nil {
		return nil, err
	}

	return dbUser, nil
}