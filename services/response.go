package services

import {
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
}

type Response struct {
	Error string  `json:"error,omitempty"`
	Result float64 `json:"result"`
}

func sendSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation %s successfull", op),
		"data":    data,
	})
}