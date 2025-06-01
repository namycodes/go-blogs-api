package blog

type Service interface {
	CreateBlog(blog *Blog) (*Blog, error)
	FindAllBlogs() ([]*Blog, error)
	FindBlogById(id uint) (*Blog, error)
	DeleteBlogById(id uint) (uint, error)
}

type serviceImpl struct {
	repo Repository
}

func NewService(repo Repository) Service{
	return &serviceImpl{repo}
}

func (s *serviceImpl) CreateBlog(blog *Blog) (*Blog, error){
	return s.repo.Create(blog)
}

func (s *serviceImpl) FindAllBlogs() ([]*Blog, error){
	return s.repo.GetAll()
}

func (s *serviceImpl) FindBlogById(id uint) (*Blog, error){
	return s.repo.GetById(id)
}

func (s *serviceImpl) DeleteBlogById(id uint) (uint, error){
	return s.repo.DeleteById(id)
}