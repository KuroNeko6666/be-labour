package models

type AuthLoginModel struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AuthRegisterModel struct {
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required" gorm:"unique"`
	Email    string `json:"email" validate:"required" gorm:"unique"`
	Password string `json:"password" validate:"required" gorm:"unique"`
}
