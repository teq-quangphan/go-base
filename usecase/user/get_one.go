package user

import "go-base/model"

func (uc *UserUseCase) GetOneUser(id int) (*model.User, error) {
	return uc.repo.GetOne(id)
}
