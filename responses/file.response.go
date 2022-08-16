package responses

import "github.com/KuroNeko6666/be-labour/models"

type FileResponse struct {
	Status  int              `json:"status"`
	Message string           `json:"message"`
	Data    models.FileModel `json:"data"`
}
