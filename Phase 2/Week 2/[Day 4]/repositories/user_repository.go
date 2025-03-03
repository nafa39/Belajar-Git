package repositories

import (
	"gorm/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(DB *gorm.DB) UserRepository {
	return &userRepository{DB}
}

func (r *userRepository) Create(user *models.User) error {
	result := r.db.Create(user)
	return result.Error
}

func (r *userRepository) FindByID(id int) (*models.User, error) {
	user := models.User{}
	result := r.db.First(user, id)
	return &user, result.Error
}
