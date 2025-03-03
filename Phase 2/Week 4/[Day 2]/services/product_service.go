package services

import (
	"errors"
	"testing/entity"
	"testing/repo"
)

type ProductService struct {
	RepoProduct repo.ProductRepo
}

func (ps ProductService) GetProductById(id int) (*entity.Product, error) {
	product := ps.RepoProduct.FindById(id)
	if product == nil {
		return nil, errors.New("product not found")
	}
	return product, nil
}
