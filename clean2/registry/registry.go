package registry

import (
	"clean2/adapters"

	"github.com/jinzhu/gorm"
)

type registry struct {
	db *gorm.DB
}

type Registry interface {
	NewAppController() adapters.AppController
}

func NewRegistry(db *gorm.DB) Registry {
	return &registry{db}
}

func (r *registry) NewAppController() adapters.AppController {
	return adapters.AppController{
		User: r.NewUserController(),
	}
}
