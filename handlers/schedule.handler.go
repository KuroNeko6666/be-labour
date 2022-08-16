package handlers

import (
	"fmt"
	"strings"

	"github.com/KuroNeko6666/be-labour/database"
	"github.com/KuroNeko6666/be-labour/models"
	"github.com/KuroNeko6666/be-labour/responses"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func CreateSchedule(c *fiber.Ctx) error {
	var schedule models.ScheduleModel

	tokenStr := strings.Split(c.GetReqHeaders()["Authorization"], " ")[1]
	token, _ := jwt.Parse(tokenStr, nil)
	claims := token.Claims.(jwt.MapClaims)
	claimAdmin := fmt.Sprintf("%v", claims["admin"])

	if claimAdmin == "false" {
		return createError(c, BAD_REQUEST, ERROR, ERR_ACCESS)
	}

	if err := c.BodyParser(&schedule); err != nil {
		return createError(c, BAD_REQUEST, ERROR, err.Error())
	}

	if err := validate.Struct(&schedule); err != nil {
		return createError(c, BAD_REQUEST, ERROR, err.Error())
	}

	database.Database.Db.Create(&schedule)

	return c.Status(OK).JSON(
		responses.ScheduleResponse{
			Status:  OK,
			Message: SUCCESS,
			Data:    schedule,
		},
	)
}

func GetSchedules(c *fiber.Ctx) error {
	schedule := []models.ScheduleModel{}

	database.Database.Db.Find(&schedule)

	return c.Status(OK).JSON(
		responses.ScheduleResponses{
			Status:  OK,
			Message: SUCCESS,
			Data:    schedule,
		},
	)

}

func GetSchedule(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var schedule models.ScheduleModel

	if err != nil {
		return createError(c, BAD_REQUEST, ERROR, ERR_ENSURE_ID)
	}

	database.Database.Db.Find(&schedule, "id = ?", id)

	if schedule.ID == 0 {
		return createError(c, NOT_FOUND, ERROR, ERR_NOT_EXIST)
	}

	return c.Status(OK).JSON(
		responses.ScheduleResponse{
			Status:  OK,
			Message: SUCCESS,
			Data:    schedule,
		},
	)

}

func UpdateSchedule(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var schedule models.ScheduleModel
	var scheduleUpdate models.ScheduleModelUpdate

	if err != nil {
		return createError(c, BAD_REQUEST, ERROR, ERR_ENSURE_ID)
	}

	if err := c.BodyParser(&scheduleUpdate); err != nil {
		return createError(c, BAD_REQUEST, ERROR, err.Error())
	}

	database.Database.Db.Find(&schedule, "id = ?", id)
	if schedule.ID == 0 {
		return createError(c, NOT_FOUND, ERROR, ERR_NOT_EXIST)
	}

	if scheduleUpdate.DateTime == "" {
		data := schedule.DateTime
		dataID := schedule.ID
		schedule = models.ScheduleModel(scheduleUpdate)
		schedule.ID = dataID
		schedule.DateTime = data
	} else if scheduleUpdate.Location == "" {
		data := schedule.Location
		dataID := schedule.ID
		schedule = models.ScheduleModel(scheduleUpdate)
		schedule.ID = dataID
		schedule.Location = data
	} else if scheduleUpdate.Room == "" {
		data := schedule.Room
		dataID := schedule.ID
		schedule = models.ScheduleModel(scheduleUpdate)
		schedule.ID = dataID
		schedule.Room = data
	} else if scheduleUpdate.Quota == 0 {
		data := schedule.Quota
		dataID := schedule.ID
		schedule = models.ScheduleModel(scheduleUpdate)
		schedule.ID = dataID
		schedule.Quota = data
	} else {
		dataID := schedule.ID
		schedule = models.ScheduleModel(scheduleUpdate)
		schedule.ID = dataID
	}

	database.Database.Db.Save(&schedule)

	return c.Status(OK).JSON(
		responses.ScheduleResponse{
			Status:  OK,
			Message: SUCCESS,
			Data:    schedule,
		},
	)
}

func DeleteSchedule(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var schedule models.ScheduleModel

	if err != nil {
		return createError(c, BAD_REQUEST, ERROR, ERR_ENSURE_ID)
	}

	database.Database.Db.Find(&schedule, "id = ?", id)
	if schedule.ID == 0 {
		return createError(c, NOT_FOUND, ERROR, ERR_NOT_EXIST)
	}

	database.Database.Db.Delete(&schedule)

	return c.Status(OK).JSON(
		responses.ResponseText{
			Status:  OK,
			Message: SUCCESS,
			Data:    SCS_DELETE,
		},
	)
}
