package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/nibroos/bookman-go/user-service/internal/controller/rest"
	"github.com/nibroos/bookman-go/user-service/internal/repository"
	"github.com/nibroos/bookman-go/user-service/internal/service"
	"gorm.io/gorm"
)

func SetupAuthorRoutes(authors fiber.Router, gormDB *gorm.DB, sqlDB *sqlx.DB) {
	authorRepo := repository.NewAuthorRepository(gormDB, sqlDB)
	authorService := service.NewAuthorService(authorRepo)
	authorController := rest.NewAuthorController(authorService)

	// prefix /authors

	authors.Post("/index-author", authorController.GetAuthors)
	authors.Post("/show-author", authorController.GetAuthorByID)
	authors.Post("/create-author", authorController.CreateAuthor)
	authors.Post("/update-author", authorController.UpdateAuthor)
	authors.Post("/delete-author", authorController.DeleteAuthor)
	authors.Post("/restore-author", authorController.RestoreAuthor)
}
