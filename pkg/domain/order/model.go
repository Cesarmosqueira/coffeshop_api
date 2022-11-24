package order

import (
	"time"

	o "github.com/Cesarmosqueira/coffeshop_api/pkg/domain/product"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID		 primitive.ObjectID		`bson:"_id" json:"id" mongo:"-"`
	CreatedBy		string			`mongo:"created_by" json:"created_by" bson:"created_by"`
	CreatedAt		*time.Time		`mongo:"created_at" json:"created_at" bson:"created_at"`
	Invoice			float32			`mongo:"invoice" json:"invoice" bson:"invoice"`
	Items			[]o.Product		`mongo:"items" json:"items" bson:"items"`
}

// returns total order price and error
func (o *Order) addItem(product o.Product) (float64, error) {
	if len(o.Items) == 0 {
		o.Invoice = 0.0
	}

	o.Items = append(o.Items, product)
	o.Invoice += product.Price

	return float64(o.Invoice), nil
}
