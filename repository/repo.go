package repository

import (
	"go-base/repository/user"
	"gorm.io/gorm"
)

type Repository struct {
	RepoUser user.IRepoUser
}

func New(db *gorm.DB) *Repository {
	return &Repository{
		RepoUser: user.NewRepo(db),
	}
}
