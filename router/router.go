package router

import (
	"github.com/gin-gonic/gin"
)
func InitRouter(){
	// Initialize Router
	router := gin.Default()

	//Initialize routes
	initializeRoutes(router)

	// Server API
	router.Run(":9000")
}