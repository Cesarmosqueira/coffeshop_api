package product

import (
	"context"

	"github.com/Cesarmosqueira/coffeshop_api/pkg/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)
type store struct {
	collection *mongo.Collection
	ctx        context.Context
}

type ProductStore interface {
	Create (Product) (Product, error)
	GetById (string) (Product, error)
	ListAll () ([]Product, error)
}

func NewProductStore() ProductStore {
	return &store{
		collection: database.DBClient.Collection("products"),
		ctx:        context.Background(),
	}
}


func (s *store) Create(product Product) (Product, error) {
	product.ID = primitive.NewObjectID()
	_, err := s.collection.InsertOne(context.TODO(), product)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (s *store) GetById(id string) (Product, error) {
	doc := s.collection.FindOne(context.TODO(), bson.M{})

	var product Product
	doc.Decode(&product)
	return product, nil


}


func (s *store) ListAll() ([]Product, error) {
	cursor, err := s.collection.Find(s.ctx, bson.D{{}})

	var products []Product
	if err = cursor.All(context.TODO(), &products); err != nil {
		panic(err)
	}
	return products, nil
}
