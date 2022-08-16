package main

import (
	"log"
	"net/http"

	"github.com/KuroNeko6666/be-labour/database"
	"github.com/KuroNeko6666/be-labour/responses"
	"github.com/KuroNeko6666/be-labour/routes"
	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(
		responses.ResponseText{
			Status:  http.StatusOK,
			Message: "success",
			Data:    "welcome to labour api",
		},
	)
}

func notFound(c *fiber.Ctx) error {
	return c.Status(http.StatusNotFound).JSON(
		responses.ResponseText{
			Status:  http.StatusNotFound,
			Message: "error",
			Data:    "url not found",
		},
	)
}

func main() {
	database.ConnectDB()
	app := fiber.New()
	app.Get("/api", welcome)
	routes.AuthRoute(app)
	routes.UserRoutes(app)
	routes.AdminRoutes(app)
	routes.CompanyRoutes(app)
	routes.ScheduleRoute(app)
	app.Use("/", notFound)

	log.Fatal(app.Listen(":8000"))
}
