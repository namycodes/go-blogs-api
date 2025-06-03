package middleware

import (
	"net/http"
	"strings"
	"com.namycodes/helpers"
	"github.com/gin-gonic/gin"
)

func JwtAuth() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader,"Bearer"){
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized"})
			return
		}
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := helpers.VerifyJwtToken(tokenStr)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		ctx.Set("user_id", claims.UserId)
		ctx.Next()
	}
	
}


