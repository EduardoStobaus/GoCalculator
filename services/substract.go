package services

import (
	"net/http"

	"github.com/EduardoStobaus/GoCalculator/models"
	"github.com/gin-gonic/gin"
)

func OperationSubstract(ctx *gin.Context) {
	req := models.CalcRequest{}

	ctx.BindJSON(&req)

	resp := models.Response{}
	if req.Op == "-" {
		resp.Result = req.FirstNumber - req.SecondNumber
		models.SendSuccess(ctx, "substraction", resp.Result)
	} else {
		models.SendError(ctx, http.StatusBadRequest, "operator defined incorrectly")
	}
}
