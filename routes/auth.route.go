package routes

import (
	"github.com/KuroNeko6666/be-labour/handlers"
	"github.com/gofiber/fiber/v2"
)

func AuthRoute(app *fiber.App) {
	app.Post("/api/login", handlers.Login)
	app.Post("/api/register", handlers.Register)
	app.Post("/api/loginAdmin", handlers.LoginAdmin)
	app.Post("/api/registerAdmin", handlers.RegisterAdmin)
	app.Post("/api/file", handlers.UploadFile)
	app.Get("/api/file/:id", handlers.DownloadFile)

}
