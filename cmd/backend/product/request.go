package product

import (
	val "github.com/Cesarmosqueira/coffeshop_api/internal/validation"
)

type ProductDto struct {
	Name	string		`json:"name"`
	Cost	float32		`json:"cost"`
}

	
func (p *ProductDto) Validate() []val.ValidationError {
	validator := val.NewValidator(*p)
	validator.NotBlank("Name")
	validator.NotBlank("Cost")

	validator.MinLength("Name", 8)

	return validator.Summary()
}
