package fetch_products

import (
	"github.com/learning-go/challenges/challenges_parser/internal/domain/product"
	"github.com/learning-go/challenges/challenges_parser/internal/infrastructure/product/repository"
)

type GetProductsService struct {
	repository repository.ProductRepository
}

func NewGetProductsService(repository repository.ProductRepository) GetProductsService {
	return GetProductsService{repository: repository}
}

func (service *GetProductsService) Execute() (*product.Products, error) {
	_products, err := service.repository.FetchAll()
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	return _products, nil
}
