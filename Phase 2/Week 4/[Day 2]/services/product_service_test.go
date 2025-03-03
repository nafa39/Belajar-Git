package services

import (
	"testing/repo"

	"github.com/stretchr/testify/mock"
)

var productRepoMock = &repo.ProductRepoMock{Mock: mock.Mock{}}
var productService = ProductService{RepoProduct: productRepoMock}

func GetProductById(id int) (*repo.Product, error) {
	return productService.GetProductById(id)
}
