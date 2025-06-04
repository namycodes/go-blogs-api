package blog

import (
	"net/http"

	"com.namycodes/helpers"
	dto "com.namycodes/internal/dto/responsedto"
	"com.namycodes/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service}
}

func (h *Handler) CreateBlog(ctx *gin.Context) {
	var blog models.Blog

	if err := ctx.ShouldBindJSON(&blog); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	createdBlog, err := h.service.CreateBlog(&blog)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	helpers.Response(ctx, "Creeated Blog Successfully", http.StatusCreated, "blog", dto.BlogResponseDto{
		Id:          createdBlog.Id,
		Title:       createdBlog.Title,
		Description: createdBlog.Description,
		UserId:      createdBlog.UserId,
	})

}

func (h *Handler) GetAllBlogs(ctx *gin.Context) {
	blogs, err := h.service.FindAllBlogs()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
	}

	var blogsResponse []dto.BlogResponseDto

	for _, b := range blogs {
		blogsResponse = append(blogsResponse, dto.BlogResponseDto{
			Id:          b.Id,
			Title:       b.Title,
			Description: b.Description,
			UserId:      b.UserId,
		})
	}

	helpers.Response(ctx, "All Blogs", http.StatusOK, "blogs", blogsResponse)
}

func (h *Handler) GetBlogById(ctx *gin.Context) {

	id := ctx.Param("id")

	blog, err := h.service.FindBlogById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	helpers.Response(ctx, "Fetched Blog Successfully", http.StatusOK, "blog", dto.BlogResponseDto{
		Id:          blog.Id,
		Title:       blog.Title,
		Description: blog.Description,
		UserId:      blog.UserId,
	})

}

func (h *Handler) DeleteBlogById(ctx *gin.Context) {
	id := ctx.Param("id")

	val, err := h.service.DeleteBlogById(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	if val == 0 {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": "Could not delete blog"})
		return
	}

	helpers.Response(ctx, "Deleted blog Successfully", http.StatusOK, "blog", nil)

}

func (h *Handler) UpdateBlog(ctx *gin.Context) {
	var updatedBlogBody models.Blog

	if err := ctx.ShouldBindJSON(&updatedBlogBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id := ctx.Param("id")

	isBlogUpdated, err := h.service.UpdateBlogById(id, &models.Blog{
		Title:       updatedBlogBody.Title,
		Description: updatedBlogBody.Description,
	})

	if err != nil || !isBlogUpdated {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	blogId, err := uuid.Parse(id)

	helpers.Response(ctx, "Updated Blog Successfully", http.StatusOK, "blog", dto.BlogResponseDto{
		Id:          blogId,
		Title:       updatedBlogBody.Title,
		Description: updatedBlogBody.Description,
	})

}
