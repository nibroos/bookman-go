package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/nibroos/bookman-go/user-service/internal/controller/rest"
	"github.com/nibroos/bookman-go/user-service/internal/repository"
	"github.com/nibroos/bookman-go/user-service/internal/service"
	"gorm.io/gorm"
)

func SetupBookRoutes(books fiber.Router, gormDB *gorm.DB, sqlDB *sqlx.DB) {
	bookRepo := repository.NewBookRepository(gormDB, sqlDB)
	bookService := service.NewBookService(bookRepo)
	bookController := rest.NewBookController(bookService)

	// prefix /books

	books.Post("/index-book", bookController.GetBooks)
	books.Post("/show-book", bookController.GetBookByID)
	books.Post("/create-book", bookController.CreateBook)
	books.Post("/update-book", bookController.UpdateBook)
	books.Post("/delete-book", bookController.DeleteBook)
	books.Post("/restore-book", bookController.RestoreBook)
}
