package order

import (
	"time"
	val "github.com/Cesarmosqueira/coffeshop_api/internal/validation"
)

type OrderDto struct {
	ID			string			`json:"id"`
	CreatedBy	string			`json:"created_by"`
	CreatedAt	*time.Time		`json:"created_at"`
	Items		[]string		`json:"items"`

}

	
func (p *OrderDto) Validate() []val.ValidationError {
	validator := val.NewValidator(*p)
	validator.NotBlank("CreatedBy")
	validator.NotBlank("CreatedAt")
	validator.NotBlank("Items")

	return validator.Summary()
}
