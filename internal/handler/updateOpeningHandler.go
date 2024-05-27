package handler

import (
	"fmt"
	"net/http"

	"github.com/WandersonLontra/gopportunities/internal/schema"
	"github.com/gin-gonic/gin"
)

func UpdateOpeningHandler(ctx *gin.Context){
	id := ctx.Query("id")

	var request UpdateOpeningDTO

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation fields error: %v",err.Error())
		sendError(ctx,http.StatusBadRequest, err.Error())
		return
	}
	foundOpening := schema.Opening{}

	if err := db.First(&foundOpening,id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening not found: %v",err))
		return
	}

	opening := schema.Opening{}

	if request.Role != "" {
		opening.Role = request.Role
	}
	if request.Company != "" {
		opening.Company = request.Company
	}
	if request.Link != "" {
		opening.Link= request.Link
	}
	if request.Location != "" {
		opening.Location = request.Location
	}
	if request.Remote != nil {
		opening.Remote = *request.Remote
	}
	if request.Salary > 0 {
		opening.Salary =  request.Salary
	}
 
	if err := db.Model(&foundOpening).Updates(&opening).Error; err != nil {
		logger.Errorf("error on update opening: %v",err.Error())
		sendError(ctx,http.StatusInternalServerError, "error on update opening in database")
		return
	}
	sendSuccess(ctx,"update",opening)
}