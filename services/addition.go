package services

import (
	"github.com/EduardoStobaus/GoCalculator/controllers"
	"github.com/EduardoStobaus/GoCalculator/schemas"
	"github.com/gin-gonic/gin"
	"fmt"
)

func OperationAdd(ctx *gin.Context) {
	req := &schemas.CalcRequest{}
	
	ctx.BindJSON(&req)

	resp := &controllers.Response{}
	if req.Operador == "+" {
		resp.Resultado = req.PrimeiroNumero + req.SegundoNumero
	} else {
		resp.Erro = fmt.Sprintf("Bad Request: Operador não definido corretamente: %s", req.Operador)
	}

}
