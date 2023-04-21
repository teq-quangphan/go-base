package user

import (
	"go-base/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type RepoUser struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) IRepoUser {
	return &RepoUser{db: db}
}

type IRepoUser interface {
	CreateUser(user *model.User) error
	GetOneUserByEmail(email string) (*model.User, error)
}

func (r *RepoUser) CreateUser(user *model.User) error {
	return r.db.Model(&model.User{}).Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "email"}}, UpdateAll: true}).
		Create(&user).Error
}
func (r *RepoUser) GetOneUserByEmail(email string) (*model.User, error) {
	user := &model.User{}
	if err := r.db.Model(&model.User{}).Where("email = ?", email).Take(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
