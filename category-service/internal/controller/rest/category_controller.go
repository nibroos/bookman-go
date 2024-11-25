package rest

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nibroos/bookman-go/user-service/internal/dtos"
	"github.com/nibroos/bookman-go/user-service/internal/middleware"
	"github.com/nibroos/bookman-go/user-service/internal/models"
	"github.com/nibroos/bookman-go/user-service/internal/service"
	"github.com/nibroos/bookman-go/user-service/internal/utils"
	"github.com/nibroos/bookman-go/user-service/internal/validators/form_requests"
)

type CategoryController struct {
	service *service.CategoryService
}

func NewCategoryController(service *service.CategoryService) *CategoryController {
	return &CategoryController{service: service}
}

func (c *CategoryController) GetCategories(ctx *fiber.Ctx) error {
	filters, ok := ctx.Locals("filters").(map[string]string)
	if !ok {
		return utils.SendResponse(ctx, utils.WrapResponse(nil, nil, "Invalid filters", http.StatusBadRequest), http.StatusBadRequest)
	}

	categories, total, err := c.service.GetCategories(ctx.Context(), filters)
	if err != nil {
		return utils.SendResponse(ctx, utils.WrapResponse(nil, nil, err.Error(), http.StatusInternalServerError), http.StatusInternalServerError)
	}

	paginationMeta := utils.CreatePaginationMeta(filters, total)

	return utils.GetResponse(ctx, categories, paginationMeta, "Categories fetched successfully", http.StatusOK, nil, nil)
}
func (c *CategoryController) CreateCategory(ctx *fiber.Ctx) error {
	var req dtos.CreateCategoryRequest

	// Use the utility function to parse the request body
	if err := utils.BodyParserWithNull(ctx, &req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"errors": err.Error(), "message": "Invalid request", "status": http.StatusBadRequest})
	}

	// Validate the request
	reqValidator := form_requests.NewCategoryStoreRequest().Validate(&req, ctx.Context())
	if reqValidator != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"errors": reqValidator, "message": "Validation failed", "status": http.StatusBadRequest})
	}

	// Extract user ID from JWT
	claims, err := middleware.GetAuthUser(ctx)
	if err != nil {
		return utils.GetResponse(ctx, nil, nil, "Unauthorized", http.StatusUnauthorized, err.Error(), nil)
	}
	userID := uint(claims["user_id"].(float64))

	createdAt := time.Now()

	author := models.Category{
		Name:        req.Name,
		Description: req.Description,
		CreatedByID: &userID,
		CreatedAt:   &createdAt,
	}

	createdCategory, err := c.service.CreateCategory(ctx.Context(), &author)
	if err != nil {
		return utils.GetResponse(ctx, nil, nil, "Failed to create author", http.StatusInternalServerError, err.Error(), nil)
	}

	params := &dtos.GetCategoryParams{ID: createdCategory.ID}
	getCategory, err := c.service.GetCategoryByID(ctx.Context(), params)
	if err != nil {
		return utils.GetResponse(ctx, nil, nil, "Category not found", http.StatusNotFound, err.Error(), nil)
	}

	filters := ctx.Locals("filters").(map[string]string)
	paginationMeta := utils.CreatePaginationMeta(filters, 1)

	return utils.GetResponse(ctx, []interface{}{getCategory}, paginationMeta, "Category created successfully", http.StatusCreated, nil, nil)
}

func (c *CategoryController) GetCategoryByID(ctx *fiber.Ctx) error {
	var req dtos.GetCategoryByIDRequest

	if err := ctx.BodyParser(&req); err != nil {
		return utils.GetResponse(ctx, nil, nil, "Category not found", http.StatusBadRequest, err.Error(), nil)
	}

	if req.ID == 0 {
		return utils.GetResponse(ctx, nil, nil, "Category not found", http.StatusBadRequest, "ID is required", nil)
	}

	params := &dtos.GetCategoryParams{ID: req.ID}
	author, err := c.service.GetCategoryByID(ctx.Context(), params)
	if err != nil {
		return utils.GetResponse(ctx, nil, nil, "Category not found", http.StatusNotFound, err.Error(), nil)
	}

	authorArray := []interface{}{author}

	filters := ctx.Locals("filters").(map[string]string)
	paginationMeta := utils.CreatePaginationMeta(filters, 1)

	return utils.GetResponse(ctx, authorArray, paginationMeta, "Category fetched successfully", http.StatusOK, nil, nil)
}

// update author
func (c *CategoryController) UpdateCategory(ctx *fiber.Ctx) error {
	var req dtos.UpdateCategoryRequest

	if err := utils.BodyParserWithNull(ctx, &req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"errors": err.Error(), "message": "Invalid request", "status": http.StatusBadRequest})
	}

	// Validate the request
	reqValidator := form_requests.NewCategoryUpdateRequest().Validate(&req, ctx.Context())
	if reqValidator != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"errors": reqValidator, "message": "Validation failed", "status": http.StatusBadRequest})
	}

	params := &dtos.GetCategoryParams{ID: req.ID}
	// Fetch the existing author to get the current data
	existingCategory, err := c.service.GetCategoryByID(ctx.Context(), params)
	if err != nil {
		return utils.GetResponse(ctx, nil, nil, "Category not found", http.StatusNotFound, err.Error(), nil)
	}

	// Extract user ID from JWT
	claims, err := middleware.GetAuthUser(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"errors": err.Error(), "message": "Unauthorized", "status": fiber.StatusUnauthorized})
	}
	userID := uint(claims["user_id"].(float64))

	author := models.Category{
		ID:          req.ID,
		Name:        req.Name,
		Description: req.Description,
		CreatedByID: &existingCategory.CreatedByID,
		UpdatedByID: &userID,
		CreatedAt:   existingCategory.CreatedAt,
	}

	updatedCategory, err := c.service.UpdateCategory(ctx.Context(), &author)
	if err != nil {
		return utils.GetResponse(ctx, nil, nil, "Failed to update author", http.StatusInternalServerError, err.Error(), nil)
	}

	params = &dtos.GetCategoryParams{ID: updatedCategory.ID}
	getCategory, err := c.service.GetCategoryByID(ctx.Context(), params)
	if err != nil {
		return utils.GetResponse(ctx, nil, nil, "Category not found", http.StatusNotFound, err.Error(), nil)
	}

	filters := ctx.Locals("filters").(map[string]string)
	paginationMeta := utils.CreatePaginationMeta(filters, 1)

	return utils.GetResponse(ctx, []interface{}{getCategory}, paginationMeta, "Category updated successfully", http.StatusOK, nil, nil)
}

// delete author
func (c *CategoryController) DeleteCategory(ctx *fiber.Ctx) error {
	var req dtos.DeleteCategoryRequest

	if err := ctx.BodyParser(&req); err != nil {
		return utils.GetResponse(ctx, nil, nil, "Category not found", http.StatusBadRequest, err.Error(), nil)
	}

	if req.ID == 0 {
		return utils.GetResponse(ctx, nil, nil, "Category not found", http.StatusBadRequest, "ID is required", nil)
	}

	params := &dtos.GetCategoryParams{ID: req.ID}
	// GET author by ID
	_, err := c.service.GetCategoryByID(ctx.Context(), params)
	if err != nil {
		return utils.GetResponse(ctx, nil, nil, "Category not found", http.StatusNotFound, err.Error(), nil)
	}

	err = c.service.DeleteCategory(ctx.Context(), req.ID)
	if err != nil {
		return utils.GetResponse(ctx, nil, nil, "Failed to delete author", http.StatusInternalServerError, err.Error(), nil)
	}

	return utils.GetResponse(ctx, nil, nil, "Category deleted successfully", http.StatusOK, nil, nil)
}

// restore author
func (c *CategoryController) RestoreCategory(ctx *fiber.Ctx) error {
	var req dtos.DeleteCategoryRequest

	if err := ctx.BodyParser(&req); err != nil {
		return utils.GetResponse(ctx, nil, nil, "Category not found", http.StatusBadRequest, err.Error(), nil)
	}

	if req.ID == 0 {
		return utils.GetResponse(ctx, nil, nil, "Category not found", http.StatusBadRequest, "ID is required", nil)
	}

	isDeleted := 1
	params := &dtos.GetCategoryParams{ID: req.ID, IsDeleted: &isDeleted}
	// GET author by ID
	_, err := c.service.GetCategoryByID(ctx.Context(), params)
	if err != nil {
		return utils.GetResponse(ctx, nil, nil, "Category not found", http.StatusNotFound, err.Error(), nil)
	}

	err = c.service.RestoreCategory(ctx.Context(), req.ID)
	if err != nil {
		return utils.GetResponse(ctx, nil, nil, "Failed to restore author", http.StatusInternalServerError, err.Error(), nil)
	}

	return utils.GetResponse(ctx, nil, nil, "Category restored successfully", http.StatusOK, nil, nil)
}
