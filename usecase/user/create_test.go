package user

import (
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go-base/model"
	"go-base/repository/user/mocks"
	"gorm.io/gorm"
	"testing"
)

func TestUserUseCase_Create(t *testing.T) {
	mockRepo := mocks.NewIRepoUser(t)
	u := UserUseCase{mockRepo}

	t.Run("success", func(t *testing.T) {
		req := model.User{
			UserName: "quangpn",
			Password: "1",
			Email:    "quangphan@gmail.com",
		}
		mockRepo.On("GetOneUserByEmail", req.Email).Return(nil, gorm.ErrRecordNotFound).Once()
		mockRepo.On("CreateUser", mock.Anything).Return(nil)

		user, err := u.Create(req)
		require.NoError(t, err)
		require.NotNil(t, user)

		mockRepo.AssertExpectations(t)
	})

	t.Run("fail_duplicate_email", func(t *testing.T) {
		req := model.User{
			UserName: "quangpn",
			Password: "1",
			Email:    "quangphan@gmail.com",
		}

		mockGetUserByEmail := &model.User{
			UserName: "quangpn24",
			Password: "$2a$14$37pZ/fAPiK9zjTNwpx8nqeg3uvcpj2.aPT3B1mD3cIWN1A4ISh1Wu",
			Email:    req.Email,
		}
		mockRepo.On("GetOneUserByEmail", req.Email).Return(mockGetUserByEmail, nil).Once()

		user, err := u.Create(req)
		require.Error(t, err)
		require.Nil(t, user)
		require.True(t, mockRepo.AssertExpectations(t))
	})
}
