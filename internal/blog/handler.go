package blog

import (
	"net/http"
	"strconv"

	"com.namycodes/internal/models"
	"github.com/gin-gonic/gin"
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

	ctx.JSON(http.StatusCreated, gin.H{"blog": createdBlog})
}

func (h *Handler) GetAllBlogs(ctx *gin.Context) {
	blogs, err := h.service.FindAllBlogs()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"blogs": blogs})
}

func (h *Handler) GetBlogById(ctx *gin.Context) {

	idParam := ctx.Param("id")

	idUint64, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	id := uint(idUint64)

	blog, err := h.service.FindBlogById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"blog": blog})
}


func (h *Handler) DeleteBlogById(ctx *gin.Context){
	idParam := ctx.Param("id")

	idUint64, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	id := uint(idUint64)
    id, err = h.service.DeleteBlogById(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"id": id})
}