package repo

import "go-test-7/entity"

type ProductRepo interface {
	FindByID(id int) *entity.Product
	// CreateProduct() *entity.Product
	// FindAll() *entity.Product
	// UpdateProduct(id int) *entity.Product
	// DeleteProduct(id int) *entity.Product
}
