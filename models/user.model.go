package models

type UserModel struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" gorm:"unique" validate:"required"`
	Email    string `json:"email" gorm:"unique" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserModelUpdate struct {
	Name     string `json:"name"`
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}
