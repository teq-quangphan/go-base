package user

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"go-base/util"
	"net/http"
	"time"
)

type RefreshTokenReq struct {
	Token string `json:"token"`
}

func (r *Route) RefreshToken(c echo.Context) error {
	refreshTokenReq := &RefreshTokenReq{}
	if err := c.Bind(&refreshTokenReq); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Error when parse req: %s", err.Error()))
	}

	//validate token
	token, err := jwt.Parse(refreshTokenReq.Token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, echo.NewHTTPError(http.StatusForbidden, "Unexpected signing method: %v", token.Header["alg"])
		}
		signature := []byte(util.JWT_SECRET_KEY)
		return signature, nil
	})

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		//check expired time
		if int64(claims[util.EXP_AT].(float64)) < time.Now().Local().Unix() {
			return c.JSON(http.StatusUnauthorized, "refresh token expired")
		}

		//create new token
		accessToken, err := util.GenerateToken(claims["id"].(string), claims["email"].(string))
		if err != nil {
			return c.JSON(http.StatusForbidden, err.Error())
		}

		res := util.Response{
			Code:    201,
			Message: "Successfully",
			Data: map[string]interface{}{
				"access_token":  accessToken,
				"refresh_token": refreshTokenReq.Token,
			},
		}
		return c.JSON(http.StatusCreated, res)
	} else {
		return c.JSON(http.StatusUnauthorized, "refresh invalid")
	}
}
