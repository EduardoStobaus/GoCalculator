package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func OperationAdd(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Addition done successfully",
	})
}
