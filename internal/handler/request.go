package handler

import "fmt"

type RequestOpeningDTO struct{
	Role 		string		`json:"role"`
	Company		string		`json:"company"`
	Location	string		`json:"location"`
	Remote		*bool		`json:"remote"`
	Link		string		`json:"link"`
	Salary		float64		`json:"salary"`
}

func (co *RequestOpeningDTO) Validate() error {
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

type UpdateOpeningDTO struct {
	Role 		string		`json:"role"`
	Company		string		`json:"company"`
	Location	string		`json:"location"`
	Remote		*bool		`json:"remote"`
	Link		string		`json:"link"`
	Salary		float64		`json:"salary"`
}

func (co *UpdateOpeningDTO) Validate() error {
	if co.Role != "" || co.Company != "" || co.Location != "" || co.Link != "" || co.Remote != nil  || co.Salary > 0{
		return nil
	}

	return fmt.Errorf("at least one field must be provided")
}