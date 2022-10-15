package usecase

import "clean2/domain/model"

type UserPresenter interface {
	ResponseUsers(u []*model.User) []*model.User
}
