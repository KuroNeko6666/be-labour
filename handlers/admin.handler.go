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

func CreateAdmin(c *fiber.Ctx) error {
	var user models.AdminModel
	tokenStr := strings.Split(c.GetReqHeaders()["Authorization"], " ")[1]
	token, _ := jwt.Parse(tokenStr, nil)
	claims := token.Claims.(jwt.MapClaims)
	claimAdmin := fmt.Sprintf("%v", claims["admin"])

	if claimAdmin == "false" {
		return createError(c, BAD_REQUEST, ERROR, ERR_ACCESS)
	}

	if err := c.BodyParser(&user); err != nil {
		return createError(c, BAD_REQUEST, ERROR, err.Error())
	}

	if err := validate.Struct(&user); err != nil {
		return createError(c, BAD_REQUEST, ERROR, err.Error())
	}

	hasher, err := createHash(user.Password)
	if err != nil {
		return createError(c, SERVER_ERROR, ERROR, ERR_SERVER_ERROR)
	}

	user.Password = hasher
	database.Database.Db.Create(&user)

	if user.ID == 0 {
		return createError(c, BAD_REQUEST, ERROR, ERR_EXIST)
	}

	responseUser := CreateResponseAdmin(user)

	return c.Status(OK).JSON(
		responses.ResponseUser{
			Status:  OK,
			Message: SUCCESS,
			Data:    responseUser,
		},
	)

}

func GetAdmins(c *fiber.Ctx) error {
	users := []models.AdminModel{}
	tokenStr := strings.Split(c.GetReqHeaders()["Authorization"], " ")[1]
	token, _ := jwt.Parse(tokenStr, nil)
	claims := token.Claims.(jwt.MapClaims)
	claimAdmin := fmt.Sprintf("%v", claims["admin"])

	if claimAdmin == "false" {
		return createError(c, BAD_REQUEST, ERROR, ERR_ACCESS)
	}

	database.Database.Db.Find(&users)

	responseUsers := []responses.Users{}

	for _, user := range users {
		responseUser := CreateResponseAdmins(user)
		responseUsers = append(responseUsers, responseUser)
	}

	return c.Status(OK).JSON(
		responses.ResponseUsers{
			Status:  OK,
			Message: SUCCESS,
			Data:    responseUsers,
		},
	)
}

func GetAdmin(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var user models.AdminModel
	tokenStr := strings.Split(c.GetReqHeaders()["Authorization"], " ")[1]
	token, _ := jwt.Parse(tokenStr, nil)
	claims := token.Claims.(jwt.MapClaims)
	claimAdmin := fmt.Sprintf("%v", claims["admin"])

	if claimAdmin == "false" {
		return createError(c, BAD_REQUEST, ERROR, ERR_ACCESS)
	}

	if err != nil {
		return createError(c, BAD_REQUEST, ERROR, ERR_ENSURE_ID)
	}

	if err := FindAdmin(id, &user); err != nil {
		return createError(c, NOT_FOUND, ERROR, ERR_NOT_EXIST)
	}

	responseUser := CreateResponseAdmin(user)

	return c.Status(OK).JSON(
		responses.ResponseUser{
			Status:  OK,
			Message: SUCCESS,
			Data:    responseUser,
		},
	)

}

func UpdateAdmin(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var user models.AdminModel
	var updateUser models.AdminModelUpdate
	tokenStr := strings.Split(c.GetReqHeaders()["Authorization"], " ")[1]
	token, _ := jwt.Parse(tokenStr, nil)
	claims := token.Claims.(jwt.MapClaims)
	claimAdmin := fmt.Sprintf("%v", claims["admin"])

	if claimAdmin == "false" {
		return createError(c, BAD_REQUEST, ERROR, ERR_ACCESS)
	}

	if err != nil {
		return createError(c, BAD_REQUEST, ERROR, ERR_ENSURE_ID)
	}

	if err := c.BodyParser(&updateUser); err != nil {
		return createError(c, BAD_REQUEST, ERROR, err.Error())
	}

	if err := FindAdmin(id, &user); err != nil {
		return createError(c, NOT_FOUND, ERROR, ERR_NOT_EXIST)
	}

	if updateUser.Password != "" {
		hasher, err := createHash(updateUser.Password)
		if err != nil {
			return createError(c, SERVER_ERROR, ERROR, ERR_SERVER_ERROR)
		}

		updateUser.Password = hasher
	}

	if updateUser.Email == "" {
		user.Name = updateUser.Name
		user.Password = updateUser.Password
		user.Username = updateUser.Username
	} else if updateUser.Username == "" {
		user.Name = updateUser.Name
		user.Password = updateUser.Password
		user.Email = updateUser.Email
	} else if updateUser.Name == "" {
		user.Email = updateUser.Email
		user.Password = updateUser.Password
		user.Username = updateUser.Username
	} else if updateUser.Password == "" {
		user.Email = updateUser.Email
		user.Email = updateUser.Email
		user.Username = updateUser.Username
	} else if updateUser.Email == user.Email {
		user.Name = updateUser.Name
		user.Password = updateUser.Password
		user.Username = updateUser.Username
	} else if updateUser.Username == user.Username {
		user.Name = updateUser.Name
		user.Password = updateUser.Password
		user.Email = updateUser.Email
	} else {
		user.Name = updateUser.Name
		user.Email = updateUser.Email
		user.Username = updateUser.Username
		user.Password = updateUser.Password
	}

	if err := database.Database.Db.Save(&user).Error; err != nil {
		return createError(c, BAD_REQUEST, ERROR, err.Error())
	}

	responseUser := CreateResponseAdmin(user)

	return c.Status(OK).JSON(
		responses.ResponseUser{
			Status:  OK,
			Message: SUCCESS,
			Data:    responseUser,
		},
	)

}

func DeleteAdmin(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var user models.AdminModel
	tokenStr := strings.Split(c.GetReqHeaders()["Authorization"], " ")[1]
	token, _ := jwt.Parse(tokenStr, nil)
	claims := token.Claims.(jwt.MapClaims)
	claimAdmin := fmt.Sprintf("%v", claims["admin"])

	if claimAdmin == "false" {
		return createError(c, BAD_REQUEST, ERROR, ERR_ACCESS)
	}

	if err != nil {
		return createError(c, BAD_REQUEST, ERROR, ERR_ENSURE_ID)
	}

	if err := FindAdmin(id, &user); err != nil {
		return createError(c, BAD_REQUEST, ERROR, ERR_NOT_FOUND)
	}

	if err := database.Database.Db.Delete(&user).Error; err != nil {
		return createError(c, NOT_FOUND, ERROR, err.Error())
	}

	return c.Status(OK).JSON(
		responses.ResponseText{
			Status:  OK,
			Message: SUCCESS,
			Data:    "Succesfully Delete User",
		},
	)
}
