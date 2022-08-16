package routes

import (
	"github.com/KuroNeko6666/be-labour/handlers"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func UserRoutes(app *fiber.App) {
	app.Use(jwtware.New(jwtware.Config{
		SigningKey:   []byte("secret"),
		ErrorHandler: handlers.TokenAware,
	}))
	app.Get("/api/users", handlers.GetUsers)
	app.Get("/api/users/:id", handlers.GetUser)
	app.Post("/api/users", handlers.CreateUser)
	app.Put("/api/users/:id", handlers.UpdateUser)
	app.Delete("/api/users/:id", handlers.DeleteUser)
}
