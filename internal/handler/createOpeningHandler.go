package handler

import (
	"fmt"
	"net/http"

	"github.com/WandersonLontra/gopportunities/internal/schema"
	"github.com/gin-gonic/gin"
)



type CreateOpeningDTO struct{
	Role 		string		`json:"role"`
	Company		string		`json:"company"`
	Location	string		`json:"location"`
	Remote		*bool		`json:"remote"`
	Link		string		`json:"link"`
	Salary		float64		`json:"salary"`
}

func (co *CreateOpeningDTO) Validate() error {
	if co.Role == "" && co.Company == "" && co.Location == "" && co.Link == "" && co.Remote == nil  && co.Salary <= 0{
		return fmt.Errorf("bad formed body data")
	}
	if co.Role == "" {
		return paramsRequiredHelper("Role","string")
	}
	if co.Company == "" {
		return paramsRequiredHelper("Company","string")
	}
	if co.Location == "" {
		return paramsRequiredHelper("Location","string")
	}
	if co.Link == "" {
		return paramsRequiredHelper("Link","string")
	}
	if co.Remote == nil {
		return paramsRequiredHelper("Remote","boolean")
	}
	if co.Salary <= 0 {
		return paramsRequiredHelper("Salary","float")
	}
	return nil
}

func paramsRequiredHelper(name, typ string) error{
	return fmt.Errorf("params: %s (type: %s) is required", name, typ)
}

func CreateOpeningHandler(ctx *gin.Context){
	request := CreateOpeningDTO{}

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