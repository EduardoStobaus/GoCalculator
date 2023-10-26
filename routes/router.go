package routes

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {
	router := gin.Default()
	port := os.Getenv("PORT")

	initializeRoutes(router)
	router.Run(port)
}
