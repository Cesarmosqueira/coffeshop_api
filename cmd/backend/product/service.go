package product

import (
	"errors"
	"log"

	p "github.com/Cesarmosqueira/coffeshop_api/pkg/domain/product"
)

type service struct {
	productStore p.ProductStore
}

type ProductService interface {
	CreateProduct(ProductDto) (p.Product, error)
	GetProduct(string) (p.Product, error)
	ListProducts() ([]p.Product, error)
}

func NewProductService() ProductService {
	return &service{
		productStore: p.NewProductStore(),
	}
}


func (s *service) CreateProduct(request ProductDto) (p.Product, error) {
	product := p.Product{
		Name:	request.Name,
		Cost:	request.Cost,
	}


	product, err := s.productStore.Create(product)
	if err != nil {
		log.Println(err)
		return product, errors.New("No se pudo crear el producto")
	}

	return product, nil
}

func (s *service) ListProducts() ([]p.Product, error) {
	products, err := s.productStore.ListAll()
	if err != nil {
		return []p.Product{}, errors.New("Could not retrieve products.")
	}

	return products, nil
}


func (s *service) GetProduct(id string) (p.Product, error) {
	product, err := s.productStore.GetById(id)
	if err != nil {
		return p.Product{}, err
	}
	return product, err;

}
