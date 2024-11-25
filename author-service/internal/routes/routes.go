package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/nibroos/bookman-go/user-service/internal/controller/rest"
	"github.com/nibroos/bookman-go/user-service/internal/middleware"
	"gorm.io/gorm"
)

// SetupRoutes sets up the REST routes for the user service.
func SetupRoutes(app *fiber.App, gormDB *gorm.DB, sqlDB *sqlx.DB) {
	// Public routes
	app.Get("/api/v1/author/test", func(c *fiber.Ctx) error {
		return c.SendString("REST Author Service!")
	})

	version := app.Group("/api/v1")

	version.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Service is running",
		})
	})

	// Protected routes
	app.Use(middleware.JWTMiddleware())

	authors := version.Group("/authors")
	SetupAuthorRoutes(authors, gormDB, sqlDB)

	// Seeder route
	version.Post("/seeders/run", rest.NewSeederController(sqlDB.DB).RunSeeders)
}
