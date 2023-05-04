package model

import (
	"github.com/google/uuid"
)

type RefreshToken struct {
	ID        uuid.UUID `json:"id"`
	Token     string    `json:"token"`
	IsExpired bool      `json:"is_expired"`
	UserId    int       `json:"user_id"`
}

func (RefreshToken) TableName() string {
	return "refresh_token"
}
