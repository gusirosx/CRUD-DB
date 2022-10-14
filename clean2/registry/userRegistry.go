package registry

import (
	"clean2/adapters"
	"clean2/usecase"
)

func (r *registry) NewUserController() adapters.UserController {
	return adapters.NewUserController(r.NewUserInteractor())
}

func (r *registry) NewUserInteractor() usecase.UserInteractor {
	return usecase.NewUserInteractor(r.NewUserRepository(), r.NewUserPresenter())
}

func (r *registry) NewUserRepository() usecase.UserRepository {
	return adapters.NewUserRepository(r.db)
}

func (r *registry) NewUserPresenter() usecase.UserPresenter {
	return adapters.NewUserPresenter()
}
