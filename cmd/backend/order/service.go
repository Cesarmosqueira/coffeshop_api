package order

import (
	"errors"
	"log"

	o "github.com/Cesarmosqueira/coffeshop_api/pkg/domain/order"
	p "github.com/Cesarmosqueira/coffeshop_api/pkg/domain/product"
)

type service struct {
	orderStore o.OrderStore
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
	}
}


func (s *service) CreateOrder(request OrderDto) (o.Order, error) {
	order := o.Order{
		CreatedBy : request.CreatedBy,
		CreatedAt : request.CreatedAt,
		Invoice	  : request.Invoice,
		Items	  : make([]p.Product, 0),
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
	return order, err;

}

func (s *service) DeleteOrder(id string) (int64, error) {
	count, err := s.orderStore.DeleteById(id)

	if err != nil {
		return count, err
	}
	return count, err;

}