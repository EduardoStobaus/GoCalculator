package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func OperationDivision(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Division done successfully",
	})
}
