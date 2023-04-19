package mysql

import (
	"gorm.io/gorm"
)

type RepoMysql struct {
	db *gorm.DB
}

func NewRepoMysql(db *gorm.DB) IRepoMysql {
	return &RepoMysql{db: db}
}

type IRepoMysql interface {
}
