package responses

type Account struct {
	Name string `json:"name"`
}

type User struct {
	ID       uint    `json:"id" gorm:"primaryKey"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Account  Account `json:"account"`
}

type Users struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type ResponseUser struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    User   `json:"data"`
}

type ResponseUsers struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Data    []Users `json:"data"`
}
