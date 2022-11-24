package order

import (
	"time"

	pr "github.com/Cesarmosqueira/coffeshop_api/cmd/backend/product"
	val "github.com/Cesarmosqueira/coffeshop_api/internal/validation"
)

type OrderDto struct {
	ID			string			`json:"id"`
	CreatedBy	string			`json:"created_by"`
	CreatedAt	*time.Time		`json:"created_at"`
	Invoice		float32			`json:"invoice"`
	Items		[]pr.ProductDto	`json:"items"`

}

	
func (p *OrderDto) Validate() []val.ValidationError {
	validator := val.NewValidator(*p)
	validator.NotBlank("CreatedBy")
	validator.NotBlank("CreatedAt")
	validator.NotBlank("Invoice")
	validator.NotBlank("Items")

	return validator.Summary()
}
