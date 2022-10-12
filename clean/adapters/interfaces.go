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
	GetUser(userID string) (*domain.User, error)
}

func NewUserUsecase(userRepo usecases.UserRepo) UserUsecase {
	return &UserUsecaseImpl{userRepo}
}

func (e *UserUsecaseImpl) GetUsers() ([]*domain.User, error) {
	return e.userRepo.GetUsers()
}

func (e *UserUsecaseImpl) GetUser(userID string) (*domain.User, error) {
	return e.userRepo.GetUser(userID)
}

// ==================================

type UserInteractor interface {
	GetUsers() ([]domain.User, error)
	GetUser() (domain.User, error)
	//Add(userId, orderId, itemId int) error
}

type ServiceHandler struct {
	UserInteractor UserInteractor
}

type UserHandler struct {
	personUsecase UserUsecase
}
