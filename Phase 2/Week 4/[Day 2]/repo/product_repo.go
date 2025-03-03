package repo

import "testing/entity"

type ProductRepo interface {
	FindById(id int) *entity.Product
	// CreateProduct() *entity.Product
}
