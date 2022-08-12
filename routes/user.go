package routes

import (
	"errors"
	"net/http"

	"github.com/KuroNeko6666/be-labour/database"
	"github.com/KuroNeko6666/be-labour/models"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	// this is not the model user, see this as the serializer
	ID        uint   `json:"id" gorm:"primaryKey"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func CreateResponseUser(userModel models.User) User {
	return User{
		ID:        userModel.ID,
		FirstName: userModel.FirstName,
		LastName:  userModel.LastName,
	}
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	database.Database.Db.Create(&user)

	responseUser := CreateResponseUser(user)

	return c.Status(http.StatusOK).JSON(responseUser)
}

func GetUsers(c *fiber.Ctx) error {
	users := []models.User{}

	database.Database.Db.Find(&users)

	responseUsers := []User{}

	for _, user := range users {
		responseUser := CreateResponseUser(user)
		responseUsers = append(responseUsers, responseUser)
	}

	return c.Status(http.StatusOK).JSON(responseUsers)
}

func FindUser(id int, user *models.User) error {
	database.Database.Db.Find(&user, "id = ?", id)
	if user.ID == 0 {
		return errors.New("user does not exist")
	}
	return nil
}

func GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("Please ensure that :id is an integer")
	}

	if err := FindUser(id, &user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	responseUser := CreateResponseUser(user)

	return c.Status(http.StatusOK).JSON(responseUser)

}

func UpdateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("Please ensure that :id is an integer")
	}

	if err := FindUser(id, &user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	type UpdateUser struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}

	var updateData UpdateUser

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	if updateData.FirstName == "" {
		user.LastName = updateData.LastName
	} else if updateData.LastName == "" {
		user.FirstName = updateData.FirstName
	} else {
		user.FirstName = updateData.FirstName
		user.LastName = updateData.LastName
	}

	database.Database.Db.Save(&user)

	responseUser := CreateResponseUser(user)
	return c.Status(http.StatusOK).JSON(responseUser)

}

func DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("Please ensure that :id is an integer")
	}

	if err := FindUser(id, &user); err != nil {
		return c.Status(http.StatusNotFound).JSON(err.Error())
	}

	if err := database.Database.Db.Delete(&user).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(err.Error())
	}

	return c.Status(http.StatusOK).SendString("Succesfully Delete User")

}
