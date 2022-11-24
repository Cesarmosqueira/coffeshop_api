package product

import (
	"context"
	"log"

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
	GetByProductCode(string) (Product, error)
	ListAll () ([]Product, error)
	DeleteById (string) (int64, error)
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
	var product Product

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil{
		log.Println("Invalid id")
		return product, err
	}

	doc := s.collection.FindOne(s.ctx, bson.M{"_id": objectId})
	doc.Decode(&product)
	return product, nil
}

func (s *store) GetByProductCode(productCode string) (Product, error) {
	var product Product

	doc := s.collection.FindOne(s.ctx, bson.M{"productCode": productCode})
	doc.Decode(&product)
	return product, nil
}
func (s *store) DeleteById(id string) (int64, error) {

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil{
		log.Println("Invalid id")
		return 0, err
	}

	doc, err := s.collection.DeleteOne(s.ctx, bson.M{"_id": objectId})
	return doc.DeletedCount, nil
}

func (s *store) ListAll() ([]Product, error) {
	cursor, err := s.collection.Find(s.ctx, bson.D{{}})

	var products []Product
	if err = cursor.All(context.TODO(), &products); err != nil {
		panic(err)
	}
	return products, nil
}

