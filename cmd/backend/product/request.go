package product

import (
	val "github.com/Cesarmosqueira/coffeshop_api/internal/validation"
)

type ProductDto struct {
	ID			string		`json:"id"`
	Name		string		`json:"name"`
	Price		float32		`json:"price"`
	Description	string		`json:"description"`
	Branch		string		`json:"branch"`
	Stars		int32		`json:"stars"`
	ImageUrl	string		`json:"imageUrl"`
}

	
func (p *ProductDto) Validate() []val.ValidationError {
	validator := val.NewValidator(*p)
	validator.NotBlank("Name")
	validator.NotBlank("Price")
	validator.NotBlank("Description")
	validator.NotBlank("Branch")
	validator.NotBlank("Stars")

	validator.MinLength("Name", 4)

	return validator.Summary()
}
