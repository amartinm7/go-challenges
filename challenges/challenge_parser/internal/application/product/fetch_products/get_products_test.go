package fetch_products

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/learning-go/challenges/challenges_parser/internal/infrastructure/product/repository/mock"
)

func Test_get_products_service(t *testing.T) {
	mockProductRepository := new(mock.ProductRepositoryMock)
	mockProductRepository.On("FetchAll").Return(&mock.DataProducts, nil)
	service := NewGetProductsService(mockProductRepository)
	products, err := service.Execute()
	mockProductRepository.AssertCalled(t, "FetchAll")
	assert.Nil(t, err)
	assert.Equal(t, products, &mock.DataProducts)
}
