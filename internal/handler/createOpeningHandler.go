package handler

import (
	"net/http"

	"github.com/WandersonLontra/gopportunities/internal/schema"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary 	Create opening
// @Description Create a new job opening
// @Tags		Openings
// @Accept		json
// @Produce		json
// @Param		request		body		RequestOpeningDTO	true	"Request body"
// @Success		200			{object}	SuccessResponse
// @Failure		400			{object}	ErrorResponse
// @Failure		500			{object}	ErrorResponse
// @Router		/opening	[post]
func CreateOpeningHandler(ctx *gin.Context){
	request := RequestOpeningDTO{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation fields error: %v",err.Error())
		sendError(ctx,http.StatusBadRequest, err.Error())
		return
	}

	opening := schema.Opening{
		Role: request.Role,
		Company: request.Company,
		Location: request.Location,
		Remote: *request.Remote,
		Link: request.Link,
		Salary: request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("error on create opening: %v",err.Error())
		sendError(ctx,http.StatusInternalServerError, "error on create opening in database")
		return
	}
	sendSuccess(ctx,"creation",opening)
}