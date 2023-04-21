package user

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"go-base/model"
	"go-base/util"
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
		return nil, err
	}
	// check password
	match := util.CheckPasswordHash(req.Password, user.Password)
	if !match {
		return nil, fmt.Errorf("password incorrect")
	}

	//generate token
	token, err = util.GenerateToken(strconv.Itoa(user.ID), user.Email)
	if err != nil {
		return nil, err
	}

	//generate refresh token
	refreshToken, err := util.GenerateRefreshToken(strconv.Itoa(user.ID), user.Email)
	if err != nil {
		return nil, err
	}
	cookie := new(http.Cookie)
	cookie.Name = "refresh_token"
	cookie.Value = refreshToken
	ctx.SetCookie(cookie)

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
