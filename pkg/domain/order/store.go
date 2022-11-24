package order

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

type OrderStore interface {
	Create (Order) (Order, error)
	GetById (string) (Order, error)
	ListAll () ([]Order, error)
	DeleteById (string) (int64, error)
}

func NewOrderStore() OrderStore {
	return &store{
		collection: database.DBClient.Collection("orders"),
		ctx:        context.Background(),
	}
}


func (s *store) Create(order Order) (Order, error) {
	order.ID = primitive.NewObjectID()
	_, err := s.collection.InsertOne(context.TODO(), order)
	if err != nil {
		return order, err
	}
	return order, nil
}

func (s *store) GetById(id string) (Order, error) {
	var order Order

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil{
		log.Println("Invalid id")
		return order, err
	}

	doc := s.collection.FindOne(s.ctx, bson.M{"_id": objectId})
	doc.Decode(&order)
	return order, nil
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

func (s *store) ListAll() ([]Order, error) {
	cursor, err := s.collection.Find(s.ctx, bson.D{{}})

	var orders []Order
	if err = cursor.All(context.TODO(), &orders); err != nil {
		panic(err)
	}
	return orders, nil
}

