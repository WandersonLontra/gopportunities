package router

import (
	docs "github.com/WandersonLontra/gopportunities/docs"
	"github.com/WandersonLontra/gopportunities/internal/handler"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine){
	// Initialize handler
	handler.InitializeHandler()
	basePath := "/api/v1"
	v1 := router.Group(basePath)
	{
		v1.GET("/opening", handler.GetOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.GET("/openings", handler.ListOpeningHandler)
	}

	// Initialize Swagger
	docs.SwaggerInfo.Title = "Gopportunities Example API"
	docs.SwaggerInfo.Description = "This is a sample server for Gopportunities server."
	docs.SwaggerInfo.Version = "1.0"
	// docs.SwaggerInfo.Host = "localhost:9000"
	docs.SwaggerInfo.BasePath = basePath
	docs.SwaggerInfo.Schemes = []string{"http"}


	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}