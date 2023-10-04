package routes

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {
	router := gin.Default()

	CalcHandleRequest()

	router.Run(":8080")
}
