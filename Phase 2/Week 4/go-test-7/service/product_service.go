package service

import (
	"errors"
	"go-test-7/entity"
	"go-test-7/repo"
)

type ProductService struct {
	RepoProduct repo.ProductRepo
}

func (ps *ProductService) GetProduct(id int) (*entity.Product, error) {
	product := ps.RepoProduct.FindByID(id)
	if product == nil {
		return nil, errors.New("product not found")
	}
	return product, nil
}
