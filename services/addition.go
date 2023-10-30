package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func OperationAdd(ctx *gin.Context) {
	req := CalcRequest{}

	ctx.BindJSON(&req)

	resp := Response{}
	if req.Op == "+" {
		resp.Result = req.FirstNumber + req.SecondNumber
		sendSuccess(ctx, "addition", resp.Result)
	} else {
		sendError(ctx, http.StatusBadRequest, "operator defined incorrectly")
	}
}
