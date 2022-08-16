package models

type CompanyModel struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" validate:"required"`
	Username    string `json:"username" gorm:"unique" validate:"required"`
	Email       string `json:"email" gorm:"unique" validate:"required"`
	Password    string `json:"password" validate:"required"`
	Address     string `json:"address"`
	Member      string `json:"member"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Vision      string `json:"vision"`
	Mision      string `json:"mision"`
}

type CompanyModelUpdate struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Username    string `json:"username" gorm:"unique"`
	Email       string `json:"email" gorm:"unique"`
	Password    string `json:"password"`
	Address     string `json:"address"`
	Member      string `json:"member"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Vision      string `json:"vision"`
	Mision      string `json:"mision"`
}
