package model

type User struct {
	BaseModel
	UserName string `json:"user_name" gorm:"not null"`
	Password string `json:"password" gorm:"not null"`
}
