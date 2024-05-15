package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
func InitRouter(){
	router := gin.Default()

	router.GET("/ping",func (c *gin.Context){
		c.JSON(http.StatusOK,gin.H{
			"message":"pong",
		})
	})
	router.Run(":9000")
}