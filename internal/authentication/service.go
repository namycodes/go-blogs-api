package authentication

import (
	"com.namycodes/helpers"
	dto "com.namycodes/internal/dto/requestdto"
	"com.namycodes/internal/models"
	"github.com/google/uuid"
)


type Service interface {
	RegisterUser(user *dto.RegisterUserDto) (*models.User, error)
	Login(user *dto.AuthenticationDto) (*models.User, error)
}

type serviceImpl struct {
    repo Repository
}

func NewService(repo Repository) Service{
    return &serviceImpl{repo}
}

func (s *serviceImpl) RegisterUser(user *dto.RegisterUserDto) (*models.User, error){
	hashedPassword, err := helpers.HashPassword(user.Password)
	userId := uuid.New()
	if err != nil {
		return nil, err
	}
	return s.repo.Create(
		&models.User{
		Id: userId,
		FirstName: user.FirstName,
		LastName: user.LastName,
		Email: user.Email,
		Gender: user.Gender,
		Password: string(hashedPassword),
	})
}

func (s *serviceImpl) Login(user *dto.AuthenticationDto) (*models.User, error){
	return s.repo.Login(user)
}