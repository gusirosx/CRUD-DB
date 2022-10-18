package usecase

import "clean2/domain/model"

type UserRepository interface {
	FindAll() ([]*model.User, error)
	FindByID(UID string) (*model.User, error)
	Create(*model.User) (*model.User, error)
}
