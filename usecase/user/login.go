package user

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go-base/model"
	"go-base/util"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func (uc *UserUseCase) Login(ctx echo.Context, req model.LoginReq) (*util.Response, error) {
	var (
		user  *model.User
		res   *util.Response
		err   error
		token string
	)

	//get user
	user, err = uc.GetOneByEmail(req.Email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ctx.JSON(http.StatusBadRequest, err.Error())
		}
		return nil, ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	// check password
	match := util.CheckPasswordHash(req.Password, user.Password)
	if !match {
		return nil, ctx.JSON(http.StatusBadRequest, "password incorrect")
	}

	//generate token
	token, err = util.GenerateToken(strconv.Itoa(user.ID), user.Email)
	if err != nil {
		return nil, ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	//generate refresh token
	refreshToken, err := util.GenerateRefreshToken(strconv.Itoa(user.ID), user.Email)
	if err != nil {
		return nil, ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	rt := &model.RefreshToken{ID: uuid.New(), Token: refreshToken, IsExpired: false, UserId: user.ID}
	if err := uc.repo.CreateRefreshToken(rt); err != nil {
		return nil, ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	//set value for res
	dataRes := model.LoginRes{
		Id:           user.ID,
		Email:        user.Email,
		UserName:     user.UserName,
		AccessToken:  token,
		RefreshToken: refreshToken,
	}
	res = &util.Response{
		Code:    200,
		Message: "Successfully!",
		Data:    dataRes,
	}

	return res, nil

}
