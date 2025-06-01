package blog

import "gorm.io/gorm"


type Repository interface {
	Create(blog *Blog) (*Blog, error)
	GetAll() ([]*Blog, error)
	GetById(id uint) (*Blog, error)
	DeleteById(id uint) (uint, error)
}

type repositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository{
	return &repositoryImpl{db: db}
}

func (r *repositoryImpl) Create(blog *Blog) (*Blog, error){
    if err:= r.db.Create(blog).Error; err != nil {
		return nil, err
	}
    
	return blog, nil
}

func (r *repositoryImpl) GetAll() ([]*Blog, error){
	var blogs []*Blog
	if err:= r.db.Find(&blogs).Error; err != nil {
		return nil, err
	}

	return blogs, nil
}

func (r *repositoryImpl) GetById(id uint) (*Blog, error){
	var blog *Blog
	if err := r.db.First(&blog, id).Error; err != nil {
		return nil, err
	}
	return blog, nil
}

func (r *repositoryImpl) DeleteById(id uint) (uint, error){
	var blog *Blog
	if err := r.db.First(&blog, id).Error; err != nil {
		return 0, nil
	}

	if err := r.db.Delete(&blog, id).Error; err != nil {
		return 0, err
	}

	return id, nil
}