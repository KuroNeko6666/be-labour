package handlers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/KuroNeko6666/be-labour/database"
	"github.com/KuroNeko6666/be-labour/models"
	"github.com/KuroNeko6666/be-labour/responses"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func CreateCompany(c *fiber.Ctx) error {
	var user models.CompanyModel

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

	responseUser := CreateResponseCompany(user)

	return c.Status(OK).JSON(
		responses.ResponseCompany{
			Status:  OK,
			Message: SUCCESS,
			Data:    responseUser,
		},
	)

}

func GetCompanies(c *fiber.Ctx) error {
	users := []models.CompanyModel{}

	database.Database.Db.Find(&users)

	responseUsers := []responses.Companies{}

	for _, user := range users {
		responseUser := CreateResponseCompanies(user)
		responseUsers = append(responseUsers, responseUser)
	}

	return c.Status(OK).JSON(
		responses.ResponseCompanies{
			Status:  OK,
			Message: SUCCESS,
			Data:    responseUsers,
		},
	)
}

func GetCompany(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var user models.CompanyModel

	if err != nil {
		return createError(c, BAD_REQUEST, ERROR, ERR_ENSURE_ID)
	}

	if err := FindCompany(id, &user); err != nil {
		return createError(c, NOT_FOUND, ERROR, ERR_NOT_EXIST)
	}

	responseUser := CreateResponseCompany(user)

	return c.Status(OK).JSON(
		responses.ResponseCompany{
			Status:  OK,
			Message: SUCCESS,
			Data:    responseUser,
		},
	)

}

func UpdateCompany(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var user models.CompanyModel
	var updateUser models.CompanyModelUpdate
	header := strings.Split(c.GetReqHeaders()["Authorization"], " ")
	tokenStr := header[1]

	if err != nil {
		return createError(c, BAD_REQUEST, ERROR, ERR_ENSURE_ID)
	}

	if err := c.BodyParser(&updateUser); err != nil {
		return createError(c, BAD_REQUEST, ERROR, err.Error())
	}

	if err := FindCompany(id, &user); err != nil {
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
		data := user.Email
		dataID := user.ID
		user = models.CompanyModel(updateUser)
		user.ID = dataID
		user.Username = data
	} else if updateUser.Username == "" {
		data := user.Username
		dataID := user.ID
		user = models.CompanyModel(updateUser)
		user.ID = dataID
		user.Username = data
	} else if updateUser.Name == "" {
		data := user.Name
		dataID := user.ID
		user = models.CompanyModel(updateUser)
		user.ID = dataID
		user.Username = data
	} else if updateUser.Password == "" {
		data := user.Password
		dataID := user.ID
		user = models.CompanyModel(updateUser)
		user.ID = dataID
		user.Username = data
	} else if updateUser.Address == "" {
		data := user.Address
		dataID := user.ID
		user = models.CompanyModel(updateUser)
		user.ID = dataID
		user.Username = data
	} else if updateUser.Description == "" {
		data := user.Description
		dataID := user.ID
		user = models.CompanyModel(updateUser)
		user.ID = dataID
		user.Username = data
	} else {
		dataID := user.ID
		user = models.CompanyModel(updateUser)
		user.ID = dataID
	}

	token, _ := jwt.Parse(tokenStr, nil)
	claims := token.Claims.(jwt.MapClaims)
	claimAdmin := fmt.Sprintf("%v", claims["admin"])
	claimID := fmt.Sprintf("%v", claims["id"])
	userID := strconv.FormatUint(uint64(user.ID), 10)

	fmt.Println(updateUser)

	if claimAdmin == "false" && claimID != userID {
		return createError(c, BAD_REQUEST, ERROR, ERR_ACCESS)
	}

	if err := database.Database.Db.Save(&user).Error; err != nil {
		return createError(c, BAD_REQUEST, ERROR, ERR_EXIST)
	}

	responseUser := CreateResponseCompany(user)

	return c.Status(OK).JSON(
		responses.ResponseCompany{
			Status:  OK,
			Message: SUCCESS,
			Data:    responseUser,
		},
	)

}

func DeleteCompany(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var user models.CompanyModel
	header := strings.Split(c.GetReqHeaders()["Authorization"], " ")
	tokenStr := header[1]

	fmt.Println(err)

	if err != nil {
		return createError(c, BAD_REQUEST, ERROR, ERR_ENSURE_ID)
	}

	if err := FindCompany(id, &user); err != nil {
		return createError(c, BAD_REQUEST, ERROR, ERR_NOT_FOUND)
	}

	token, _ := jwt.Parse(tokenStr, nil)
	claims := token.Claims.(jwt.MapClaims)
	claimAdmin := fmt.Sprintf("%v", claims["admin"])
	claimID := fmt.Sprintf("%v", claims["id"])
	userID := strconv.FormatUint(uint64(user.ID), 10)

	fmt.Println(userID, claimID)

	if claimAdmin == "false" && claimID != userID {
		return createError(c, BAD_REQUEST, ERROR, ERR_ACCESS)
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
