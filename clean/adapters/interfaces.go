package adapters

import (
	"CRUD-DB/clean/domain"
	"CRUD-DB/clean/usecases"
)

type UserUsecaseImpl struct {
	userRepo usecases.UserRepo
}

type UserUsecase interface {
	GetUsers() ([]*domain.User, error)
}

func NewUserUsecase(userRepo usecases.UserRepo) UserUsecase {
	return &UserUsecaseImpl{userRepo}
}

func (e *UserUsecaseImpl) GetUsers() ([]*domain.User, error) {
	return e.userRepo.GetUsers()
}
