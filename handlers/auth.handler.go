package handlers

import (
	"github.com/KuroNeko6666/be-labour/database"
	"github.com/KuroNeko6666/be-labour/models"
	"github.com/KuroNeko6666/be-labour/responses"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *fiber.Ctx) error {
	var userLogin models.AuthLoginModel
	var user models.UserModel

	if err := c.BodyParser(&userLogin); err != nil {
		return createError(c, BAD_REQUEST, ERROR, err.Error())
	}

	if err := validate.Struct(&userLogin); err != nil {
		return createError(c, BAD_REQUEST, ERROR, err.Error())
	}

	rows := database.Database.Db.Find(&user, "email = ?", userLogin.Email).RowsAffected

	if rows == 0 {
		return createError(c, NOT_FOUND, ERROR, ERR_NOT_FOUND)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userLogin.Password)); err != nil {
		return createError(c, BAD_REQUEST, ERROR, ERR_MATCH_PASS)
	}

	claims := jwt.MapClaims{
		"id":    user.ID,
		"admin": false,
	}

	jwt := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwt.SignedString([]byte("secret"))

	if err != nil {
		return createError(c, SERVER_ERROR, ERROR, err.Error())
	}

	responseUser := CreateResponseUser(user)

	return c.Status(OK).JSON(
		responses.LoginResponse{
			Status:  OK,
			Message: SUCCESS,
			Token:   token,
			Data:    responseUser,
		},
	)

}

func Register(c *fiber.Ctx) error {
	var userRegister models.AuthRegisterModel
	var user models.UserModel

	if err := c.BodyParser(&userRegister); err != nil {
		return createError(c, BAD_REQUEST, ERROR, err.Error())
	}

	if err := validate.Struct(&userRegister); err != nil {
		return createError(c, BAD_REQUEST, ERROR, err.Error())
	}

	hasher, err := createHash(userRegister.Password)

	if err != nil {
		return createError(c, SERVER_ERROR, ERROR, err.Error())
	}

	user.Name = userRegister.Name
	user.Username = userRegister.Username
	user.Email = userRegister.Email
	user.Password = hasher

	database.Database.Db.Create(&user)

	if user.ID == 0 {
		return createError(c, BAD_REQUEST, ERROR, ERR_EXIST)
	}

	responseUser := CreateResponseUser(user)

	return c.Status(OK).JSON(
		responses.ResponseUser{
			Status:  OK,
			Message: SUCCESS,
			Data:    responseUser,
		},
	)
}

func LoginAdmin(c *fiber.Ctx) error {
	var userLogin models.AuthLoginModel
	var user models.AdminModel

	if err := c.BodyParser(&userLogin); err != nil {
		return createError(c, BAD_REQUEST, ERROR, err.Error())
	}

	if err := validate.Struct(&userLogin); err != nil {
		return createError(c, BAD_REQUEST, ERROR, err.Error())
	}

	rows := database.Database.Db.Find(&user, "email = ?", userLogin.Email).RowsAffected

	if rows == 0 {
		return createError(c, NOT_FOUND, ERROR, ERR_NOT_FOUND)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userLogin.Password)); err != nil {
		return createError(c, BAD_REQUEST, ERROR, ERR_MATCH_PASS)
	}

	claims := jwt.MapClaims{
		"id":    user.ID,
		"admin": true,
	}

	jwt := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwt.SignedString([]byte("secret"))

	if err != nil {
		return createError(c, SERVER_ERROR, ERROR, err.Error())
	}

	responseUser := CreateResponseAdmin(user)

	return c.Status(OK).JSON(
		responses.LoginResponse{
			Status:  OK,
			Message: SUCCESS,
			Token:   token,
			Data:    responseUser,
		},
	)

}

func RegisterAdmin(c *fiber.Ctx) error {
	var userRegister models.AuthRegisterModel
	var user models.AdminModel

	if err := c.BodyParser(&userRegister); err != nil {
		return createError(c, BAD_REQUEST, ERROR, err.Error())
	}

	if err := validate.Struct(&userRegister); err != nil {
		return createError(c, BAD_REQUEST, ERROR, err.Error())
	}

	hasher, err := createHash(userRegister.Password)

	if err != nil {
		return createError(c, SERVER_ERROR, ERROR, err.Error())
	}

	user.Name = userRegister.Name
	user.Username = userRegister.Username
	user.Email = userRegister.Email
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
