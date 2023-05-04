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

//go:generate mockery --name IRepoUser
type IRepoUser interface {
	CreateUser(user *model.User) error
	GetOneUserByEmail(email string) (*model.User, error)
	CreateRefreshToken(rt *model.RefreshToken) error
	GetOne(id string) (*model.RefreshToken, error)
}

func (r *RepoUser) CreateRefreshToken(rt *model.RefreshToken) error {
	return r.db.Model(&model.RefreshToken{}).Create(&rt).Error
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
func (r *RepoUser) GetOne(id string) (*model.RefreshToken, error) {
	rt := &model.RefreshToken{}
	r.db.Model(&model.RefreshToken{}).Where("id=?", id).Take(&rt)
	return rt, nil
}
