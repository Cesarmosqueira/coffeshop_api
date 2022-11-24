package product

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID		 primitive.ObjectID `bson:"_id" json:"id" mongo:"-"`
	Name			string		`mongo:"name" json:"name" bson:"name"`
	Price			float32		`mongo:"price" json:"price" bson:"price"`
	Description		string		`mongo:"description" json:"description" bson:"description"`
	Branch			string		`mongo:"branch" json:"branch" bson:"branch"`
	Stars			int32		`mongo:"stars" json:"stars" bson:"stars"`
	ImageUrl		string		`mongo:"imageUrl" json:"imageUrl" bson:"imageUrl"`
}

