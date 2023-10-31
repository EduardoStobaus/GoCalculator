package services

import (
	"net/http"

	"github.com/EduardoStobaus/GoCalculator/models"
	"github.com/gin-gonic/gin"
)

func OperationDivision(ctx *gin.Context) {
	req := models.CalcRequest{}

	ctx.BindJSON(&req)

	if req.Op == "/" {
		resp := req.FirstNumber / req.SecondNumber
		if req.SecondNumber == 00.00 {
			models.SendError(ctx, http.StatusBadRequest, "division per 0 isn't valid")
		} else {
			models.SendSuccess(ctx, "division", resp)
		}
	} else {
		models.SendError(ctx, http.StatusBadRequest, "operator defined incorrectly")
	}
}
