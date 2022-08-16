package handlers

import (
	"fmt"

	"github.com/KuroNeko6666/be-labour/database"
	"github.com/KuroNeko6666/be-labour/models"
	"github.com/KuroNeko6666/be-labour/responses"
	"github.com/gofiber/fiber/v2"
)

func UploadFile(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	var fileModel models.FileModel

	if err != nil {
		return createError(c, BAD_REQUEST, ERROR, err.Error())
	}

	fileModel.FileName = file.Filename
	fileModel.FilePath = fmt.Sprintf("./storage/%s", file.Filename)

	database.Database.Db.Create(&fileModel)
	if fileModel.ID == 0 {
		return createError(c, BAD_REQUEST, ERROR, "file exist")
	}
	c.SaveFile(file, fmt.Sprintf("./storage/%s", file.Filename))

	return c.Status(OK).JSON(
		responses.FileResponse{
			Status:  OK,
			Message: ERROR,
			Data:    fileModel,
		},
	)
}

func DownloadFile(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var fileModel models.FileModel

	if err != nil {
		return createError(c, BAD_REQUEST, ERROR, ERR_ENSURE_ID)
	}

	database.Database.Db.Find(&fileModel, "id = ?", id)

	if fileModel.ID == 0 {
		return createError(c, NOT_FOUND, ERROR, ERR_NOT_FOUND)
	}

	return c.Status(OK).Download(fileModel.FilePath, fileModel.FileName)

}

// func DeleteFile(c *fiber.Ctx) error {
// 	id, err := c.ParamsInt("id")
// 	var fileModel models.FileModel

// 	if err != nil {
// 		return createError(c, BAD_REQUEST, ERROR, ERR_ENSURE_ID)
// 	}

// 	database.Database.Db.Find(&fileModel, "id = ?", id)

// 	if file
// }
