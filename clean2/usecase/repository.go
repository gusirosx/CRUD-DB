package usecase

import "clean2/domain/model"

type UserRepository interface {
	FindAll() ([]*model.User, error)
}
