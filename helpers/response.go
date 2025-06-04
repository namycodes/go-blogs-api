package helpers

import "github.com/gin-gonic/gin"

func Response(ctx *gin.Context,message string, statusCode int, dataLabel string, data interface{}) {
	 ctx.JSON(int(statusCode), gin.H{"messsage": message, "data": gin.H{dataLabel: data}, "statusCode": statusCode})
}