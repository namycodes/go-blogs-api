package blog

import (
	"com.namycodes/internal/models"
)

type Service interface {
	CreateBlog(blog *models.Blog) (*models.Blog, error)
	FindAllBlogs() ([]*models.Blog, error)
	FindBlogById(id string) (*models.Blog, error)
	DeleteBlogById(id string) (uint, error)
	UpdateBlogById(id string, updatedBlod *models.Blog) (bool, error)
}

type serviceImpl struct {
	repo Repository
}

func NewService(repo Repository) Service{
	return &serviceImpl{repo}
}

func (s *serviceImpl) CreateBlog(blog *models.Blog) (*models.Blog, error){
	return s.repo.Create(blog)
}

func (s *serviceImpl) FindAllBlogs() ([]*models.Blog, error){
	return s.repo.GetAll()
}

func (s *serviceImpl) FindBlogById(id string) (*models.Blog, error){
	return s.repo.GetById(id)
}

func (s *serviceImpl) DeleteBlogById(id string) (uint, error){
	return s.repo.DeleteById(id)
}

func (s *serviceImpl) UpdateBlogById(id string,updatedBlog *models.Blog) (bool, error){
	return s.repo.Update(id, updatedBlog)
}