package service

import (
	"go-test-7/entity"
	"go-test-7/repo"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var productRepoMock = &repo.ProductRepoMock{Mock: mock.Mock{}}
var productService = ProductService{RepoProduct: productRepoMock}

func TestGetProductNotFound(t *testing.T) {
	productRepoMock.Mock.On("FindByID", 1).Return(nil)

	product, err := productService.GetProduct(1)
	assert.Nil(t, product) // expected : product =  nil
	assert.NotNil(t, err)  // expected : error =  not nil
	assert.Equal(t, "product not found", err.Error(), "Test Error product not found")
}

func TestGetProduct(t *testing.T) {
	productRes := entity.Product{
		ID:   2,
		Name: "Pisang Goreng",
	}

	productRepoMock.Mock.On("FindByID", 2).Return(productRes)

	product, err := productService.GetProduct(2)
	assert.Nil(t, err)        // expected : error =  nil
	assert.NotNil(t, product) // expected : product = not nil
	assert.Equal(t, productRes.ID, product.ID, "ID should be same")
	assert.Equal(t, productRes.Name, product.Name, "Product name should be same")
	assert.Equal(t, &productRes, product, "Object product should be same")

}
