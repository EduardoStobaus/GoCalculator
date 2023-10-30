package controllers

import (
	"github.com/EduardoStobaus/GoCalculator/services"
	"github.com/gin-gonic/gin"
)

func calculatorRoutes(router *gin.Engine) {

	calcOps := router.Group("/calc")

	calcOps.GET("/add", services.OperationAdd)

	calcOps.GET("/subs", services.OperationSubstract)

	calcOps.GET("/mult", services.OperationMultiply)

	calcOps.GET("/div", services.OperationDivision)

}
