package routes

import (
	"github.com/KuroNeko6666/be-labour/handlers"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func CompanyRoutes(app *fiber.App) {
	app.Use(jwtware.New(jwtware.Config{
		SigningKey:   []byte("secret"),
		ErrorHandler: handlers.TokenAware,
	}))
	app.Get("/api/companies", handlers.GetCompanies)
	app.Get("/api/company/:id", handlers.GetCompany)
	app.Post("/api/company", handlers.CreateCompany)
	app.Put("/api/company/:id", handlers.UpdateCompany)
	app.Delete("/api/company/:id", handlers.DeleteCompany)
}
