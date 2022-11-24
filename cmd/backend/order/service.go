package order

import (
	"errors"
	"log"

	o "github.com/Cesarmosqueira/coffeshop_api/pkg/domain/order"
	p "github.com/Cesarmosqueira/coffeshop_api/pkg/domain/product"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type service struct {
	orderStore o.OrderStore
	productStore p.ProductStore
}

type OrderService interface {
	CreateOrder(OrderDto) (o.Order, error)
	GetOrder(string) (o.Order, error)
	ListOrders() ([]o.Order, error)
	DeleteOrder(string) (int64, error)
}

func NewOrderService() OrderService {
	return &service{
		orderStore: o.NewOrderStore(),
		productStore: p.NewProductStore(),
	}
}


func (s *service) CreateOrder(request OrderDto) (o.Order, error) {
	order := o.Order{
		CreatedBy : request.CreatedBy,
		CreatedAt : request.CreatedAt,
		Invoice	  : 0.0,
		Items	  : make([]p.Product, 0),
	}

	for _, productId := range request.Items {

		product, err := s.productStore.GetById(productId)
		if err != nil {
			log.Println(err)
			return order, err
		}

		order.Invoice += product.Price

		order.Items = append(order.Items, product)
	}

	order, err := s.orderStore.Create(order)
	if err != nil {
		log.Println(err)
		return order, errors.New("No se pudo crear el ordero")
	}


	return order, nil
}

func (s *service) ListOrders() ([]o.Order, error) {
	orders, err := s.orderStore.ListAll()
	if err != nil {
		return []o.Order{}, errors.New("Could not retrieve orders.")
	}

	return orders, nil
}


func (s *service) GetOrder(id string) (o.Order, error) {
	order, err := s.orderStore.GetById(id)

	if err != nil {
		return order, err
	}
	hex, _ := primitive.ObjectIDFromHex("000000000000000000000000")
	if order.ID == hex {
		return order, errors.New("'" + id + "' not found")
	}
	return order, err;

}

func (s *service) DeleteOrder(id string) (int64, error) {
	count, err := s.orderStore.DeleteById(id)

	if err != nil {
		return count, err
	}
	return count, err;

}
