package routes

import (
	"github.com/KuroNeko6666/be-labour/handlers"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func AdminRoutes(app *fiber.App) {
	app.Use(jwtware.New(jwtware.Config{
		SigningKey:   []byte("secret"),
		ErrorHandler: handlers.TokenAware,
	}))
	app.Get("/api/admins", handlers.GetAdmins)
	app.Get("/api/admin/:id", handlers.GetAdmin)
	app.Post("/api/admin", handlers.CreateAdmin)
	app.Put("/api/admin/:id", handlers.UpdateAdmin)
	app.Delete("/api/admin/:id", handlers.DeleteAdmin)
}
