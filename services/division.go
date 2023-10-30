package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func OperationDivision(ctx *gin.Context) {
	req := CalcRequest{}

	ctx.BindJSON(&req)

	resp := Response{}
	if req.Op == "/" {
		resp.Result = req.FirstNumber / req.SecondNumber
		if req.SecondNumber == 00.00 {
			sendError(ctx, http.StatusBadRequest, "division per 0 isn't valid")
		} else {
			sendSuccess(ctx, "division", resp.Result)
		}
	} else {
		sendError(ctx, http.StatusBadRequest, "operator defined incorrectly")
	}
}
