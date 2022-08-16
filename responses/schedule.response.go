package responses

import "github.com/KuroNeko6666/be-labour/models"

type ScheduleResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    models.ScheduleModel
}

type ScheduleResponses struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []models.ScheduleModel
}
