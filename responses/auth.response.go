package responses

type LoginResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    User   `json:"data"`
	Token   string `json:"token"`
}
