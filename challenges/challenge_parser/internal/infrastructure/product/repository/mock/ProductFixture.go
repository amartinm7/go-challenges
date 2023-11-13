package mock

import (
	"github.com/learning-go/challenges/challenges_parser/internal/domain/product"
)

var Iphone9, _ = product.NewProduct(1,
	"iPhone 9",
	"An apple mobile which is nothing like apple",
	549,
	12.96,
	4.69,
	94,
	"Apple",
	"smartphones",
	"https://i.dummyjson.com/data/products/1/thumbnail.jpg",
	[]string{
		"https://i.dummyjson.com/data/products/1/1.jpg",
		"https://i.dummyjson.com/data/products/1/2.jpg",
		"https://i.dummyjson.com/data/products/1/3.jpg",
		"https://i.dummyjson.com/data/products/1/4.jpg",
		"https://i.dummyjson.com/data/products/1/thumbnail.jpg",
	},
)

var IphoneX, _ = product.NewProduct(2,
	"iPhone X",
	"SIM-Free, Model A19211 6.5-inch Super Retina HD display with OLED technology A12 Bionic chip with ...",
	899,
	17.94,
	4.44,
	34,
	"Apple",
	"smartphones",
	"https://i.dummyjson.com/data/products/2/thumbnail.jpg",
	[]string{
		"https://i.dummyjson.com/data/products/2/1.jpg",
		"https://i.dummyjson.com/data/products/2/2.jpg",
		"https://i.dummyjson.com/data/products/2/3.jpg",
		"https://i.dummyjson.com/data/products/2/thumbnail.jpg",
	},
)

var DataProducts = product.NewProducts([]product.Product{*Iphone9, *IphoneX}, 2, 0, 10)

var JsonProduct = "{\"products\":[{\"id\":1,\"title\":\"iPhone 9\",\"description\":\"An apple mobile which is nothing like apple\",\"price\":549,\"discountPercentage\":12.96,\"rating\":4.69,\"stock\":94,\"brand\":\"Apple\",\"category\":\"smartphones\",\"thumbnail\":\"https://i.dummyjson.com/data/products/1/thumbnail.jpg\",\"images\":[\"https://i.dummyjson.com/data/products/1/1.jpg\",\"https://i.dummyjson.com/data/products/1/2.jpg\",\"https://i.dummyjson.com/data/products/1/3.jpg\",\"https://i.dummyjson.com/data/products/1/4.jpg\",\"https://i.dummyjson.com/data/products/1/thumbnail.jpg\"]},{\"id\":2,\"title\":\"iPhone X\",\"description\":\"SIM-Free, Model A19211 6.5-inch Super Retina HD display with OLED technology A12 Bionic chip with ...\",\"price\":899,\"discountPercentage\":17.94,\"rating\":4.44,\"stock\":34,\"brand\":\"Apple\",\"category\":\"smartphones\",\"thumbnail\":\"https://i.dummyjson.com/data/products/2/thumbnail.jpg\",\"images\":[\"https://i.dummyjson.com/data/products/2/1.jpg\",\"https://i.dummyjson.com/data/products/2/2.jpg\",\"https://i.dummyjson.com/data/products/2/3.jpg\",\"https://i.dummyjson.com/data/products/2/thumbnail.jpg\"]}],\"total\":2,\"skip\":0,\"limit\":10}"
