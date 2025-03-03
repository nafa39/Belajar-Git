package repo

import (
	"testing/entity"

	"github.com/stretchr/testify/mock"
)

type ProductRepoMock struct {
	Mock mock.Mock
}

func (m *ProductRepoMock) FindById(id int) *entity.Product {
	args := m.Mock.Called(id)
	if args.Get(0) == nil {
		return nil
	}

	product := args.Get(0).(*entity.Product)
	return product
}
