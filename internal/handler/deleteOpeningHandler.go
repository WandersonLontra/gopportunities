package handler

import (
	"fmt"
	"net/http"

	"github.com/WandersonLontra/gopportunities/internal/schema"
	"github.com/gin-gonic/gin"
)

func DeleteOpeningHandler(ctx *gin.Context){
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, paramsRequiredHelper("id","string").Error())
		return
	}
	var opening schema.Opening

	// Find opening
	if err := db.First(&opening,id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found",id))
		return
	}
	// Delete opening
	if err := db.Delete(&opening, id).Error; err != nil {
		logger.Errorf("error on delete opening: %v",err)
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error on delete opening with id: %s",id))
	}
	sendSuccess(ctx, "deletion",opening)
}
