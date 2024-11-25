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

type AuthorController struct {
	service *service.AuthorService
}

func NewAuthorController(service *service.AuthorService) *AuthorController {
	return &AuthorController{service: service}
}

func (c *AuthorController) GetAuthors(ctx *fiber.Ctx) error {
	filters, ok := ctx.Locals("filters").(map[string]string)
	if !ok {
		return utils.SendResponse(ctx, utils.WrapResponse(nil, nil, "Invalid filters", http.StatusBadRequest), http.StatusBadRequest)
	}

	authors, total, err := c.service.GetAuthors(ctx.Context(), filters)
	if err != nil {
		return utils.SendResponse(ctx, utils.WrapResponse(nil, nil, err.Error(), http.StatusInternalServerError), http.StatusInternalServerError)
	}

	paginationMeta := utils.CreatePaginationMeta(filters, total)

	return utils.GetResponse(ctx, authors, paginationMeta, "Authors fetched successfully", http.StatusOK, nil, nil)
}
func (c *AuthorController) CreateAuthor(ctx *fiber.Ctx) error {
	var req dtos.CreateAuthorRequest

	// Use the utility function to parse the request body
	if err := utils.BodyParserWithNull(ctx, &req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"errors": err.Error(), "message": "Invalid request", "status": http.StatusBadRequest})
	}

	// Validate the request
	reqValidator := form_requests.NewAuthorStoreRequest().Validate(&req, ctx.Context())
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

	author := models.Author{
		Name:        req.Name,
		Description: req.Description,
		CreatedByID: &userID,
		CreatedAt:   &createdAt,
	}

	createdAuthor, err := c.service.CreateAuthor(ctx.Context(), &author)
	if err != nil {
		return utils.GetResponse(ctx, nil, nil, "Failed to create author", http.StatusInternalServerError, err.Error(), nil)
	}

	params := &dtos.GetAuthorParams{ID: createdAuthor.ID}
	getAuthor, err := c.service.GetAuthorByID(ctx.Context(), params)
	if err != nil {
		return utils.GetResponse(ctx, nil, nil, "Author not found", http.StatusNotFound, err.Error(), nil)
	}

	filters := ctx.Locals("filters").(map[string]string)
	paginationMeta := utils.CreatePaginationMeta(filters, 1)

	return utils.GetResponse(ctx, []interface{}{getAuthor}, paginationMeta, "Author created successfully", http.StatusCreated, nil, nil)
}

func (c *AuthorController) GetAuthorByID(ctx *fiber.Ctx) error {
	var req dtos.GetAuthorByIDRequest

	if err := ctx.BodyParser(&req); err != nil {
		return utils.GetResponse(ctx, nil, nil, "Author not found", http.StatusBadRequest, err.Error(), nil)
	}

	if req.ID == 0 {
		return utils.GetResponse(ctx, nil, nil, "Author not found", http.StatusBadRequest, "ID is required", nil)
	}

	params := &dtos.GetAuthorParams{ID: req.ID}
	author, err := c.service.GetAuthorByID(ctx.Context(), params)
	if err != nil {
		return utils.GetResponse(ctx, nil, nil, "Author not found", http.StatusNotFound, err.Error(), nil)
	}

	authorArray := []interface{}{author}

	filters := ctx.Locals("filters").(map[string]string)
	paginationMeta := utils.CreatePaginationMeta(filters, 1)

	return utils.GetResponse(ctx, authorArray, paginationMeta, "Author fetched successfully", http.StatusOK, nil, nil)
}

// update author
func (c *AuthorController) UpdateAuthor(ctx *fiber.Ctx) error {
	var req dtos.UpdateAuthorRequest

	if err := utils.BodyParserWithNull(ctx, &req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"errors": err.Error(), "message": "Invalid request", "status": http.StatusBadRequest})
	}

	// Validate the request
	reqValidator := form_requests.NewAuthorUpdateRequest().Validate(&req, ctx.Context())
	if reqValidator != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"errors": reqValidator, "message": "Validation failed", "status": http.StatusBadRequest})
	}

	params := &dtos.GetAuthorParams{ID: req.ID}
	// Fetch the existing author to get the current data
	existingAuthor, err := c.service.GetAuthorByID(ctx.Context(), params)
	if err != nil {
		return utils.GetResponse(ctx, nil, nil, "Author not found", http.StatusNotFound, err.Error(), nil)
	}

	// Extract user ID from JWT
	claims, err := middleware.GetAuthUser(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"errors": err.Error(), "message": "Unauthorized", "status": fiber.StatusUnauthorized})
	}
	userID := uint(claims["user_id"].(float64))

	author := models.Author{
		ID:          req.ID,
		Name:        req.Name,
		Description: req.Description,
		CreatedByID: &existingAuthor.CreatedByID,
		UpdatedByID: &userID,
		CreatedAt:   existingAuthor.CreatedAt,
	}

	updatedAuthor, err := c.service.UpdateAuthor(ctx.Context(), &author)
	if err != nil {
		return utils.GetResponse(ctx, nil, nil, "Failed to update author", http.StatusInternalServerError, err.Error(), nil)
	}

	params = &dtos.GetAuthorParams{ID: updatedAuthor.ID}
	getAuthor, err := c.service.GetAuthorByID(ctx.Context(), params)
	if err != nil {
		return utils.GetResponse(ctx, nil, nil, "Author not found", http.StatusNotFound, err.Error(), nil)
	}

	filters := ctx.Locals("filters").(map[string]string)
	paginationMeta := utils.CreatePaginationMeta(filters, 1)

	return utils.GetResponse(ctx, []interface{}{getAuthor}, paginationMeta, "Author updated successfully", http.StatusOK, nil, nil)
}

// delete author
func (c *AuthorController) DeleteAuthor(ctx *fiber.Ctx) error {
	var req dtos.DeleteAuthorRequest

	if err := ctx.BodyParser(&req); err != nil {
		return utils.GetResponse(ctx, nil, nil, "Author not found", http.StatusBadRequest, err.Error(), nil)
	}

	if req.ID == 0 {
		return utils.GetResponse(ctx, nil, nil, "Author not found", http.StatusBadRequest, "ID is required", nil)
	}

	params := &dtos.GetAuthorParams{ID: req.ID}
	// GET author by ID
	_, err := c.service.GetAuthorByID(ctx.Context(), params)
	if err != nil {
		return utils.GetResponse(ctx, nil, nil, "Author not found", http.StatusNotFound, err.Error(), nil)
	}

	err = c.service.DeleteAuthor(ctx.Context(), req.ID)
	if err != nil {
		return utils.GetResponse(ctx, nil, nil, "Failed to delete author", http.StatusInternalServerError, err.Error(), nil)
	}

	return utils.GetResponse(ctx, nil, nil, "Author deleted successfully", http.StatusOK, nil, nil)
}

// restore author
func (c *AuthorController) RestoreAuthor(ctx *fiber.Ctx) error {
	var req dtos.DeleteAuthorRequest

	if err := ctx.BodyParser(&req); err != nil {
		return utils.GetResponse(ctx, nil, nil, "Author not found", http.StatusBadRequest, err.Error(), nil)
	}

	if req.ID == 0 {
		return utils.GetResponse(ctx, nil, nil, "Author not found", http.StatusBadRequest, "ID is required", nil)
	}

	isDeleted := 1
	params := &dtos.GetAuthorParams{ID: req.ID, IsDeleted: &isDeleted}
	// GET author by ID
	_, err := c.service.GetAuthorByID(ctx.Context(), params)
	if err != nil {
		return utils.GetResponse(ctx, nil, nil, "Author not found", http.StatusNotFound, err.Error(), nil)
	}

	err = c.service.RestoreAuthor(ctx.Context(), req.ID)
	if err != nil {
		return utils.GetResponse(ctx, nil, nil, "Failed to restore author", http.StatusInternalServerError, err.Error(), nil)
	}

	return utils.GetResponse(ctx, nil, nil, "Author restored successfully", http.StatusOK, nil, nil)
}
