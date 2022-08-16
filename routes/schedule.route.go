package routes

import (
	"github.com/KuroNeko6666/be-labour/handlers"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func ScheduleRoute(app *fiber.App) {
	app.Use(jwtware.New(jwtware.Config{
		SigningKey:   []byte("secret"),
		ErrorHandler: handlers.TokenAware,
	}))
	app.Get("/api/schedules", handlers.GetSchedules)
	app.Get("/api/schedule/:id", handlers.GetSchedule)
	app.Post("/api/schedule", handlers.CreateSchedule)
	app.Put("/api/schedule/:id", handlers.UpdateSchedule)
	app.Delete("/api/schedule/:id", handlers.DeleteSchedule)
}
