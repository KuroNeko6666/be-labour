package responses

type AccountCompany struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	Member      string `json:"member"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Vision      string `json:"vision"`
	Mision      string `json:"mision"`
}

type Company struct {
	ID       uint           `json:"id" gorm:"primaryKey"`
	Username string         `json:"username"`
	Email    string         `json:"email"`
	Account  AccountCompany `json:"account"`
}

type Companies struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	Description string `json:"description"`
}

type ResponseCompany struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Data    Company `json:"data"`
}

type ResponseCompanies struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    []Companies `json:"data"`
}
