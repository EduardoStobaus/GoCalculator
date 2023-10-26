package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func OperationSubstract(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Substract done successfully",
	})
}
