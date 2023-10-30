package services

import (
	"net/http"

	"github.com/EduardoStobaus/GoCalculator/models"
	"github.com/gin-gonic/gin"
)

func OperationMultiply(ctx *gin.Context) {
	req := models.CalcRequest{}

	ctx.BindJSON(&req)

	if req.Op == "*" {
		resp := req.FirstNumber * req.SecondNumber
		models.SendSuccess(ctx, "multiplication", resp)
	} else {
		models.SendError(ctx, http.StatusBadRequest, "operator defined incorrectly")
	}
}
