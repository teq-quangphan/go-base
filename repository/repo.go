package repository

import (
	"go-base/repository/mysql"
	"gorm.io/gorm"
)

type Repository struct {
	RepoMysql mysql.IRepoMysql
}

func New(db *gorm.DB) *Repository {
	return &Repository{
		RepoMysql: mysql.NewRepoMysql(db),
	}
}
