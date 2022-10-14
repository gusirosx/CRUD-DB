package usecase

import "clean2/domain/model"

type UserRepository interface {
	FindAll(u []*model.User) ([]*model.User, error)
}
