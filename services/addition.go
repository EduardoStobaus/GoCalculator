package services

import (
	"net/http"

	"github.com/EduardoStobaus/GoCalculator/models"
	"github.com/gin-gonic/gin"
)

func OperationAdd(ctx *gin.Context) {
	req := models.CalcRequest{}

	ctx.BindJSON(&req)

	resp := models.Response{}
	if req.Op == "+" {
		resp.Result = req.FirstNumber + req.SecondNumber
		models.SendSuccess(ctx, "addition", resp.Result)
	} else {
		models.SendError(ctx, http.StatusBadRequest, "operator defined incorrectly")
	}
}
