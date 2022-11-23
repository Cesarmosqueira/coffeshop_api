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
	GetProduct(string) (ProductDto, error)
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
		Price		:	 request.Price,
		Description	:	 request.Description,
		Branch		:	 request.Branch,
		Stars		:	 request.Stars,
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


func (s *service) GetProduct(id string) (ProductDto, error) {
	product, err := s.productStore.GetById(id)
	productDto := ProductDto{
		Name		:	product.Name,
		Price		:	 product.Price,
		Description	:	 product.Description,
		Branch		:	 product.Branch,
		Stars		:	 product.Stars,
		ImageUrl	:	 product.ImageUrl,
	}
	if err != nil {
		return ProductDto{}, err
	}
	return productDto, err;

}
