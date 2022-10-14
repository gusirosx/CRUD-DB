package adapters

import (
	"clean2/domain/model"
	"clean2/usecase"

	"github.com/jinzhu/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) usecase.UserRepository {
	return &userRepository{db}
}

func (ur *userRepository) FindAll(u []*model.User) ([]*model.User, error) {
	err := ur.db.Find(&u).Error

	if err != nil {
		return nil, err
	}

	return u, nil
}
