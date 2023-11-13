package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/learning-go/challenges/challenges_parser/internal/infrastructure/product/repository/mock"
)

//var DataProductApiResponse = ProductResponse{
//	Id:                 1,
//	Title:              "iPhone 9",
//	Description:        "An apple mobile which is nothing like apple",
//	Price:              549,
//	DiscountPercentage: 12.96,
//	Rating:             4.69,
//	Stock:              94,
//	Brand:              "Apple",
//	Category:           "smartphones",
//	Thumbnail:          "https://i.dummyjson.com/data/products/1/thumbnail.jpg",
//	Images: []string{
//		"https://i.dummyjson.com/data/products/1/1.jpg",
//		"https://i.dummyjson.com/data/products/1/2.jpg",
//		"https://i.dummyjson.com/data/products/1/3.jpg",
//		"https://i.dummyjson.com/data/products/1/4.jpg",
//		"https://i.dummyjson.com/data/products/1/thumbnail.jpg",
//	},
//}

func Test_product_repository(t *testing.T) {
	t.Run("should get all products", func(t *testing.T) {
		products, error := NewProductRepositoryImpl(&mock.HttpClientMock{}).FetchAll()
		assert.Nil(t, error)
		assert.Equal(t, *products, mock.DataProducts)
	})
}
