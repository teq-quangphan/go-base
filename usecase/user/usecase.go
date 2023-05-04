package user

import (
	"github.com/labstack/echo/v4"
	"go-base/model"
	"go-base/repository"
	"go-base/repository/user"
	"go-base/util"
)

type UserUseCase struct {
	repo user.IRepoUser
}
type IUserUseCase interface {
	Create(user model.User) (*model.User, error)
	GetOneByEmail(email string) (*model.User, error)
	Login(ctx echo.Context, req model.LoginReq) (*util.Response, error)
}

func New(repo *repository.Repository) IUserUseCase {
	return &UserUseCase{
		repo: repo.RepoUser,
	}
}
