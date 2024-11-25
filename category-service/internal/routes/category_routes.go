package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/nibroos/bookman-go/user-service/internal/controller/rest"
	"github.com/nibroos/bookman-go/user-service/internal/repository"
	"github.com/nibroos/bookman-go/user-service/internal/service"
	"gorm.io/gorm"
)

func SetupCategoryRoutes(categoriess fiber.Router, gormDB *gorm.DB, sqlDB *sqlx.DB) {
	categoriesRepo := repository.NewCategoryRepository(gormDB, sqlDB)
	categoriesService := service.NewCategoryService(categoriesRepo)
	categoriesController := rest.NewCategoryController(categoriesService)

	// prefix /categoriess

	categoriess.Post("/index-categories", categoriesController.GetCategories)
	categoriess.Post("/show-categories", categoriesController.GetCategoryByID)
	categoriess.Post("/create-categories", categoriesController.CreateCategory)
	categoriess.Post("/update-categories", categoriesController.UpdateCategory)
	categoriess.Post("/delete-categories", categoriesController.DeleteCategory)
	categoriess.Post("/restore-categories", categoriesController.RestoreCategory)
}
