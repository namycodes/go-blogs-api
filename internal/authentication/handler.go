package authentication

import (
	"net/http"

	"com.namycodes/helpers"
	dto "com.namycodes/internal/dto/requestdto"
	dto_two "com.namycodes/internal/dto/responsedto"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service}
}

func (h *Handler) RegisterUser(ctx *gin.Context) {
	var newUser dto.RegisterUserDto

	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	createdUser, err := h.service.RegisterUser(&newUser)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"user": &dto_two.RegisteredUserResponseDto{
		FirstName: createdUser.FirstName,
		LastName:  createdUser.LastName,
		Email:     createdUser.Email,
		Gender:    createdUser.Gender,
	}})
}

func (h *Handler) Login(ctx *gin.Context) {
	var user dto.AuthenticationDto

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	dbUser, err := h.service.Login(&user)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Wrong Email or Password"})
		return
	}

	isPasswordMatch, err := helpers.VerifyPassword([]byte(dbUser.Password), user.Password)
	if err != nil || !isPasswordMatch {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Wrong Email or Password"})
		return
	}

	token, err := helpers.GenerateJwtToken(dbUser.Id.String())

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Wrong Email or Password"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})

}
