package mock

import (
	"github.com/stretchr/testify/mock"

	"github.com/learning-go/challenges/challenges_parser/internal/domain/product"
)

type ProductRepositoryMock struct {
	mock.Mock
}

func NewProductRepositoryMock(mock mock.Mock) *ProductRepositoryMock {
	return &ProductRepositoryMock{Mock: mock}
}

func (m *ProductRepositoryMock) FetchAll() (*product.Products, error) {
	args := m.Called()
	return args.Get(0).(*product.Products), args.Error(1)
}
