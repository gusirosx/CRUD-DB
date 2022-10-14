package registry

import (
	"clean2/adapters"
	"clean2/adapters/controller"
	ip "clean2/adapters/presenter"
	"clean2/usecase"
)

func (r *registry) NewUserController() controller.UserController {
	return controller.NewUserController(r.NewUserInteractor())
}

func (r *registry) NewUserInteractor() usecase.UserInteractor {
	return usecase.NewUserInteractor(r.NewUserRepository(), r.NewUserPresenter())
}

func (r *registry) NewUserRepository() usecase.UserRepository {
	return adapters.NewUserRepository(r.db)
}

func (r *registry) NewUserPresenter() usecase.UserPresenter {
	return ip.NewUserPresenter()
}
