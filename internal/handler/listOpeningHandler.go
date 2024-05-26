package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/WandersonLontra/gopportunities/internal/schema"
	"github.com/gin-gonic/gin"
)

func ListOpeningHandler(ctx *gin.Context){
	var (
		limit int = 2
		offset int = 0
	)
	
	queryLimit := ctx.Query("limit")
	queryOffset := ctx.Query("offset")

	if queryLimit != "" {
		lmt, _:= strconv.Atoi(queryLimit)
		limit = lmt
	}
	if queryOffset != "" {
		off, _:= strconv.Atoi(queryOffset)
		offset = off
	}
	
	var openings []schema.Opening

	if err := db.Limit(limit).Offset(offset).Find(&openings).Error; err != nil {
		logger.Errorf("Error on list openings: %v",err)
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("error on list openings"))
		return
	}
	sendSuccess(ctx, "list",openings)
}