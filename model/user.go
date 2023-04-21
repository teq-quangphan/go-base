package model

type User struct {
	BaseModel
	UserName string `json:"user_name" validate:"required" gorm:"not null"`
	Password string `json:"password" validate:"required" gorm:"not null"`
	Email    string `json:"email" validate:"required,email"`
}

type LoginReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginRes struct {
	Id           int    `json:"id"`
	UserName     string `json:"user_name"`
	Email        string `json:"email"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
