package product

import (
	"errors"
	"log"

	p "github.com/Cesarmosqueira/coffeshop_api/pkg/domain/product"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type service struct {
	productStore p.ProductStore
}

type ProductService interface {
	CreateProduct(ProductDto) (p.Product, error)
	GetProduct(string) (p.Product, error)
	GetByProductCode(string) (p.Product, error)
	ListProducts() ([]p.Product, error)
	DeleteProduct(string) (int64, error)
}

func NewProductService() ProductService {
	return &service{
		productStore: p.NewProductStore(),
	}
}


func (s *service) CreateProduct(request ProductDto) (p.Product, error) {
	product := p.Product{
		Name:	request.Name,
		Price		:	 request.Price,
		Description	:	 request.Description,
		Branch		:	 request.Branch,
		Stars		:	 request.Stars,
		ProductCode :    request.ProductCode,
		ImageUrl	:	 request.ImageUrl,
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
		return product, err
	}
	hex, _ := primitive.ObjectIDFromHex("000000000000000000000000")
	if product.ID == hex {
		return product, errors.New("'" + id + "' not found")
	}
	return product, err;

}

func (s *service) GetByProductCode(productCode string) (p.Product, error) {
	product, err := s.productStore.GetByProductCode(productCode)

	if err != nil {
		return product, err
	}
	hex, _ := primitive.ObjectIDFromHex("000000000000000000000000")
	if product.ID == hex {
		return product, errors.New("Product code #'" +  productCode + "' not found")
	}
	return product, err;
}

func (s *service) DeleteProduct(id string) (int64, error) {
	count, err := s.productStore.DeleteById(id)

	if err != nil {
		return count, err
	}
	return count, err;

}
