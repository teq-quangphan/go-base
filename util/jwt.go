package util

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const (
	EXP_AT = "exp_at"
)

func GenerateToken(userId string, email string) (string, error) {
	claims := jwt.MapClaims{}
	claims["id"] = userId
	claims["email"] = email
	claims[EXP_AT] = time.Now().Add(EXPIRED_TIME * time.Second).Local().Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(JWT_SECRET_KEY))
}
func GenerateRefreshToken(userId string, email string) (string, error) {
	claims := jwt.MapClaims{}
	claims["id"] = userId
	claims["email"] = email
	claims[EXP_AT] = time.Now().Add(EXPIRED_TIME_REFRESH_DAYS * time.Hour * 24).Local().Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(JWT_SECRET_KEY))
}
func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func HashPassword(password string) (string, error) {
	newPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", fmt.Errorf("error when hash password, err: %s", err.Error())
	}
	return string(newPassword), nil
}
