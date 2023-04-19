package usecase

import (
	"go-base/repository"
	"go-base/usecase/user"
)

type UseCase struct {
	UserUseCase user.IUserUseCase
}

func New(repo *repository.Repository) *UseCase {
	return &UseCase{
		UserUseCase: user.New(repo),
	}
}
