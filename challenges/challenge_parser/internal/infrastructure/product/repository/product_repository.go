package repository

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/learning-go/challenges/challenges_parser/internal/domain/product"
)

type HttpClient interface {
	Get(url string) (resp *http.Response, err error)
}

type ProductResponse struct {
	Id                 int      `json:"id"`
	Title              string   `json:"title"`
	Description        string   `json:"description"`
	Price              int      `json:"price"`
	DiscountPercentage float32  `json:"discountPercentage"`
	Rating             float32  `json:"rating"`
	Stock              int      `json:"stock"`
	Brand              string   `json:"brand"`
	Category           string   `json:"category"`
	Thumbnail          string   `json:"thumbnail"`
	Images             []string `json:"images"`
}

type ProductsResponse struct {
	Items []ProductResponse `json:"products"`
	Total int               `json:"total"`
	Skip  int               `json:"skip"`
	Limit int               `json:"limit"`
}

type ProductRepository interface {
	FetchAll() (*product.Products, error)
}

type ProductRepositoryImpl struct {
	httpClient HttpClient
}

func NewProductRepositoryImpl(httpClient HttpClient) *ProductRepositoryImpl {
	return &ProductRepositoryImpl{
		httpClient: httpClient,
	}
}

func mapProductJsonToDomain(productResponse []ProductResponse) []product.Product {
	productList := make([]product.Product, 0)
	for _, item := range productResponse {
		product, _ := product.NewProduct(item.Id, item.Title, item.Description, item.Price, item.DiscountPercentage, item.Rating, item.Stock, item.Brand, item.Category, item.Thumbnail, item.Images)
		productList = append(productList, *product)
	}
	return productList
}

func mapProductsJsonToDomain(productsResponse ProductsResponse) *product.Products {
	products := product.NewProducts(
		mapProductJsonToDomain(productsResponse.Items),
		productsResponse.Total, productsResponse.Skip, productsResponse.Limit,
	)
	return &products
}

func (repository *ProductRepositoryImpl) FetchAll() (*product.Products, error) {
	var productsResponse = &ProductsResponse{}
	resp, err := repository.httpClient.Get("https://dummyjson.com/products")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(responseData, productsResponse)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(productsResponse)
	return mapProductsJsonToDomain(*productsResponse), nil
}
