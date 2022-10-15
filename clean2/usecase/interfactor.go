package usecase

import (
	"clean2/domain/model"
)

type userInteractor struct {
	UserRepository UserRepository
	UserPresenter  UserPresenter
}

type UserInteractor interface {
	Get(u []*model.User) ([]*model.User, error)
}

func NewUserInteractor(r UserRepository, p UserPresenter) UserInteractor {
	return &userInteractor{r, p}
}

func (us *userInteractor) Get(u []*model.User) ([]*model.User, error) {
	userList, err := us.UserRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return us.UserPresenter.ResponseUsers(userList), nil
}
