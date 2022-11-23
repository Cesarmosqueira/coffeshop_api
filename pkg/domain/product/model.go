package product

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID     primitive.ObjectID `bson:"_id" json:"id" mongo:"-"`
	Name	string		`mongo:"name" json:"name" bson:"name"`
	Cost	float32		`mongo:"cost" json:"cost" bson:"cost"`
}
