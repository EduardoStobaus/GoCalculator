package controllers

import (
	"os"

	"github.com/gin-gonic/gin"
)

func Initialize() {
	router := gin.Default()
	port := ":" + os.Getenv("PORT")

	calculatorRoutes(router)
	router.Run(port)
}
