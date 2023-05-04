package user

import "go-base/model"

func (uc *UserUseCase) GetOneByEmail(email string) (*model.User, error) {
	return uc.repo.GetOneUserByEmail(email)
}
