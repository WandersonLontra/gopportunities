package handler

import (
	"fmt"
	"net/http"

	"github.com/WandersonLontra/gopportunities/internal/schema"
	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code,gin.H{
		"message":msg,
	})
}

func sendSuccess(ctx *gin.Context, operation string, data interface{}) {
	ctx.Header("Content-type","application/json")
	ctx.JSON(http.StatusOK,gin.H{
		"message":fmt.Sprintf("operation %s successful",operation),
		"data": data,
	})
}


type ErrorResponse struct {
	Message string `json:"message"`
}

type SuccessResponse struct {
	Message string 					`json:"message"`
	Data	schema.OpeningResponse	`json:"data"`
}

