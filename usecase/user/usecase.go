package user

import (
	"go-base/model"
	"go-base/repository"
)

type UserUseCase struct {
	repo *repository.Repository
}
type IUserUseCase interface {
	Create(user model.User) error
}

func New(repo *repository.Repository) IUserUseCase {
	return &UserUseCase{
		repo: repo,
	}
}
