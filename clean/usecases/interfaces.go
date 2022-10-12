package usecases

import (
	"CRUD-DB/clean/domain"
	"database/sql"
)

type UserRepo interface {
	GetUsers() ([]*domain.User, error)
	GetUser(userID string) (*domain.User, error)
}

type UserRepoImpl struct {
	DB *sql.DB
}

func NewUserRepoImpl(DB *sql.DB) UserRepo {
	return &UserRepoImpl{DB}
}
