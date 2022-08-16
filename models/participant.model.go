package models

type ParticipantModels struct {
	ID          uint   `json:"id"`
	ScheduleID  int    `json:"schedule_id"`
	UserID      int    `json:"user_id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Gender      string `json:"gender"`
	NumberPhone string `json:"phone"`
	Address     string `json:"addres"`
	Status      string `json:"status"`
}
