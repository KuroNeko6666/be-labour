package main

import (
	"log"

	"github.com/KuroNeko6666/be-labour/database"
	"github.com/KuroNeko6666/be-labour/routes"
	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("welcome to my res api")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api", welcome)
	app.Get("/api/users", routes.GetUsers)
	app.Get("/api/users/:id", routes.GetUser)
	app.Post("/api/users", routes.CreateUser)
	app.Put("/api/users/:id", routes.UpdateUser)
	app.Delete("/api/users/:id", routes.DeleteUser)

}

func main() {
	database.ConnectDB()
	app := fiber.New()
	setupRoutes(app)
	app.Get("/api", welcome)

	log.Fatal(app.Listen(":3000"))
}
