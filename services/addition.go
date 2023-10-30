package services

import (
	"net/http"

	"github.com/EduardoStobaus/GoCalculator/models"
	"github.com/gin-gonic/gin"
)

func OperationAdd(ctx *gin.Context) {
	req := models.CalcRequest{}

	ctx.BindJSON(&req)

	if err := req.Validate(); err != nil {
		models.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if req.Op == "+" {
		resp := float64(req.FirstNumber) + float64(req.SecondNumber)
		models.SendSuccess(ctx, "addition", resp)
	} else {
		models.SendError(ctx, http.StatusBadRequest, "operator defined incorrectly")
	}
}
