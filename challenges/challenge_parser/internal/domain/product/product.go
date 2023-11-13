package product

import "errors"

var InvalidProductId = errors.New("invalid productId: nil")

var InvalidProductDesc = errors.New("invalid description: nil")

type Product struct {
	id                 int
	title              string
	description        string
	price              int
	discountPercentage float32
	rating             float32
	stock              int
	brand              string
	category           string
	thumbnail          string
	images             []string
}

func NewProduct(id int, title string, description string, price int, discountPercentage float32, rating float32, stock int, brand string, category string, thumbnail string, images []string) (*Product, error) {
	if id == 0 {
		return nil, InvalidProductId
	}
	if title == "" {
		return nil, InvalidProductDesc
	}
	return &Product{id: id, title: title, description: description, price: price, discountPercentage: discountPercentage, rating: rating, stock: stock, brand: brand, category: category, thumbnail: thumbnail, images: images}, nil
}

type Products struct {
	items []Product
	total int
	skip  int
	limit int
}

func NewProducts(items []Product, total int, skip int, limit int) Products {
	return Products{items: items, total: total, skip: skip, limit: limit}
}
