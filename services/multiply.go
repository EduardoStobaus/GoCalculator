package services

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func OperationMultiply(ctx *gin.Context) {
	req := CalcRequest{}

	ctx.BindJSON(&req)

	resp := Response{}
	if req.Op == "*" {
		resp.Result = req.FirstNumber * req.SecondNumber
	} else {
		resp.Error = fmt.Sprintf("Bad Request: Operador não definido corretamente: %s", req.Op)
	}

	sendSuccess(ctx, "multiplication", resp.Result)
}
