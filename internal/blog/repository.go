package blog

import (
	"database/sql"

	"com.namycodes/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	Create(blog *models.Blog) (*models.Blog, error)
	GetAll() ([]*models.Blog, error)
	GetById(id string) (*models.Blog, error)
	DeleteById(id string) (uint, error)
	Update(id string, blog *models.Blog) (bool, error)
}

type repositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repositoryImpl{db: db}
}

func (r *repositoryImpl) Create(blog *models.Blog) (*models.Blog, error) {
	var user models.User

	if err := r.db.Raw("SELECT * FROM users WHERE id = @userId", sql.Named("userId", blog.UserId)).Scan(&user).Error; err != nil {
		return nil, err
	}
	blog.Id = uuid.New()
	blog.User = user
	if err := r.db.Create(blog).Error; err != nil {
		return nil, err
	}

	return blog, nil
}

func (r *repositoryImpl) GetAll() ([]*models.Blog, error) {
	var blogs []*models.Blog
	if err := r.db.Find(&blogs).Error; err != nil {
		return nil, err
	}

	return blogs, nil
}

func (r *repositoryImpl) GetById(id string) (*models.Blog, error) {
	var blog *models.Blog
	if err := r.db.Preload("User").First(&blog, "id =? ", id).Error; err != nil {
		return nil, err
	}
	return blog, nil
}

func (r *repositoryImpl) DeleteById(id string) (uint, error) {
	var blog *models.Blog
	if err := r.db.First(&blog, "id =?", id).Error; err != nil {
		return 0, nil
	}

	if err := r.db.Delete(&blog, "id =?", id).Error; err != nil {
		return 0, err
	}

	return 1, nil
}


func (r *repositoryImpl) Update(id string, updatedBlog *models.Blog) (bool, error){
	if err := r.db.Exec("UPDATE blogs SET title = @blogTitle, description = @blogDescription WHERE id = @blogId", sql.Named("blogTitle", updatedBlog.Title), sql.Named("blogDescription", updatedBlog.Description), sql.Named("blogId", id)).Error; err != nil {
        return false, err
	}

	return true, nil
}