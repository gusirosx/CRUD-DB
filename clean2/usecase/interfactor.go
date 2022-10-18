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
	GetByID(UID string) (*model.User, error)
	Create(u *model.User) (*model.User, error)
	Update(u *model.User) (*model.User, error)
	Delete(UID string) (*model.User, error)
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

func (us *userInteractor) GetByID(UID string) (*model.User, error) {
	user, err := us.UserRepository.FindByID(UID)
	if err != nil {
		return nil, err
	}
	return us.UserPresenter.ResponseUser(user), nil
}

func (us *userInteractor) Create(user *model.User) (*model.User, error) {
	user, err := us.UserRepository.Create(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *userInteractor) Update(user *model.User) (*model.User, error) {
	user, err := us.UserRepository.Update(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *userInteractor) Delete(UID string) (*model.User, error) {
	user, err := us.UserRepository.Delete(UID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
