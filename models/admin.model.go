package models

type AdminModel struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" gorm:"unique" validate:"required"`
	Email    string `json:"email" gorm:"unique" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AdminModelUpdate struct {
	Name     string `json:"name"`
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}
