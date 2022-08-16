package handlers

import (
	"errors"

	"github.com/KuroNeko6666/be-labour/database"
	"github.com/KuroNeko6666/be-labour/models"
	"github.com/KuroNeko6666/be-labour/responses"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Welcome(c *fiber.Ctx) error {
	return c.Status(OK).JSON(
		responses.ResponseText{
			Status:  OK,
			Message: SUCCESS,
			Data:    "Welcome to labour api",
		},
	)
}

func TokenAware(c *fiber.Ctx, err error) error {
	return c.Status(401).JSON(
		responses.ResponseText{
			Status:  401,
			Message: "err",
			Data:    err.Error(),
		},
	)
}

func NotFound(c *fiber.Ctx) error {
	return c.Status(NOT_FOUND).JSON(
		responses.ResponseText{
			Status:  NOT_FOUND,
			Message: ERROR,
			Data:    "url not found",
		},
	)
}

func FindUser(id int, user *models.UserModel) error {
	database.Database.Db.Find(&user, "id = ?", id)
	if user.ID == 0 {
		return errors.New(ERR_NOT_EXIST)
	}
	return nil
}

func FindAdmin(id int, user *models.AdminModel) error {
	database.Database.Db.Find(&user, "id = ?", id)
	if user.ID == 0 {
		return errors.New(ERR_NOT_EXIST)
	}
	return nil
}

func FindCompany(id int, user *models.CompanyModel) error {
	database.Database.Db.Find(&user, "id = ?", id)
	if user.ID == 0 {
		return errors.New(ERR_NOT_EXIST)
	}
	return nil
}

func createHash(password string) (string, error) {
	byte, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(byte), err
}

func createError(c *fiber.Ctx, status int, message string, data string) error {
	return c.Status(BAD_REQUEST).JSON(
		responses.ResponseText{
			Status:  BAD_REQUEST,
			Message: ERROR,
			Data:    data,
		},
	)
}

func CreateResponseUser(userModel models.UserModel) responses.User {
	return responses.User{
		ID:       userModel.ID,
		Username: userModel.Username,
		Email:    userModel.Email,
		Account: responses.Account{
			Name: userModel.Name,
		},
	}
}

func CreateResponseUsers(userModel models.UserModel) responses.Users {
	return responses.Users{
		ID:       userModel.ID,
		Name:     userModel.Name,
		Username: userModel.Username,
		Email:    userModel.Email,
	}
}

func CreateResponseAdmin(userModel models.AdminModel) responses.User {
	return responses.User{
		ID:       userModel.ID,
		Username: userModel.Username,
		Email:    userModel.Email,
		Account: responses.Account{
			Name: userModel.Name,
		},
	}
}

func CreateResponseAdmins(userModel models.AdminModel) responses.Users {
	return responses.Users{
		ID:       userModel.ID,
		Name:     userModel.Name,
		Username: userModel.Username,
		Email:    userModel.Email,
	}
}

func CreateResponseCompany(userModel models.CompanyModel) responses.Company {
	return responses.Company{
		ID:       userModel.ID,
		Username: userModel.Username,
		Email:    userModel.Email,
		Account: responses.AccountCompany{
			Name:        userModel.Name,
			Address:     userModel.Address,
			Member:      userModel.Member,
			Description: userModel.Description,
			URL:         userModel.URL,
			Vision:      userModel.Vision,
			Mision:      userModel.Mision,
		},
	}
}

func CreateResponseCompanies(userModel models.CompanyModel) responses.Companies {
	return responses.Companies{
		ID:          userModel.ID,
		Name:        userModel.Name,
		Username:    userModel.Username,
		Email:       userModel.Email,
		Address:     userModel.Address,
		Description: userModel.Description,
	}
}
