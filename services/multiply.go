package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func OperationMultiply(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Multiply done successfully",
	})
}