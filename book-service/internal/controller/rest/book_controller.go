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

type BookController struct {
	service *service.BookService
}

func NewBookController(service *service.BookService) *BookController {
	return &BookController{service: service}
}

func (c *BookController) GetBooks(ctx *fiber.Ctx) error {
	filters, ok := ctx.Locals("filters").(map[string]string)
	if !ok {
		return utils.SendResponse(ctx, utils.WrapResponse(nil, nil, "Invalid filters", http.StatusBadRequest), http.StatusBadRequest)
	}

	books, total, err := c.service.GetBooks(ctx.Context(), filters)
	if err != nil {
		return utils.SendResponse(ctx, utils.WrapResponse(nil, nil, err.Error(), http.StatusInternalServerError), http.StatusInternalServerError)
	}

	paginationMeta := utils.CreatePaginationMeta(filters, total)

	return utils.GetResponse(ctx, books, paginationMeta, "Books fetched successfully", http.StatusOK, nil, nil)
}
func (c *BookController) CreateBook(ctx *fiber.Ctx) error {
	var req dtos.CreateBookRequest

	// Use the utility function to parse the request body
	if err := utils.BodyParserWithNull(ctx, &req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"errors": err.Error(), "message": "Invalid request", "status": http.StatusBadRequest})
	}

	// Validate the request
	reqValidator := form_requests.NewBookStoreRequest().Validate(&req, ctx.Context())
	if reqValidator != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"errors": reqValidator, "message": "Validation failed", "status": http.StatusBadRequest})
	}

	// Extract user ID from JWT
	claims, err := middleware.GetAuthUser(ctx)
	if err != nil {
		return utils.GetResponse(ctx, nil, nil, "Unbookized", http.StatusUnauthorized, err.Error(), nil)
	}
	userID := uint(claims["user_id"].(float64))

	createdAt := time.Now()

	book := models.Book{
		Name:        req.Name,
		Description: req.Description,
		CategoryID:  req.CategoryID,
		AuthorID:    req.AuthorID,
		CreatedByID: &userID,
		CreatedAt:   &createdAt,
	}

	createdBook, err := c.service.CreateBook(ctx.Context(), &book)
	if err != nil {
		return utils.GetResponse(ctx, nil, nil, "Failed to create book", http.StatusInternalServerError, err.Error(), nil)
	}

	params := &dtos.GetBookParams{ID: createdBook.ID}
	getBook, err := c.service.GetBookByID(ctx.Context(), params)
	if err != nil {
		return utils.GetResponse(ctx, nil, nil, "Book not found", http.StatusNotFound, err.Error(), nil)
	}

	filters := ctx.Locals("filters").(map[string]string)
	paginationMeta := utils.CreatePaginationMeta(filters, 1)

	return utils.GetResponse(ctx, []interface{}{getBook}, paginationMeta, "Book created successfully", http.StatusCreated, nil, nil)
}

func (c *BookController) GetBookByID(ctx *fiber.Ctx) error {
	var req dtos.GetBookByIDRequest

	if err := ctx.BodyParser(&req); err != nil {
		return utils.GetResponse(ctx, nil, nil, "Book not found", http.StatusBadRequest, err.Error(), nil)
	}

	if req.ID == 0 {
		return utils.GetResponse(ctx, nil, nil, "Book not found", http.StatusBadRequest, "ID is required", nil)
	}

	params := &dtos.GetBookParams{ID: req.ID}
	book, err := c.service.GetBookByID(ctx.Context(), params)
	if err != nil {
		return utils.GetResponse(ctx, nil, nil, "Book not found", http.StatusNotFound, err.Error(), nil)
	}

	bookArray := []interface{}{book}

	filters := ctx.Locals("filters").(map[string]string)
	paginationMeta := utils.CreatePaginationMeta(filters, 1)

	return utils.GetResponse(ctx, bookArray, paginationMeta, "Book fetched successfully", http.StatusOK, nil, nil)
}

// update book
func (c *BookController) UpdateBook(ctx *fiber.Ctx) error {
	var req dtos.UpdateBookRequest

	if err := utils.BodyParserWithNull(ctx, &req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"errors": err.Error(), "message": "Invalid request", "status": http.StatusBadRequest})
	}

	// Validate the request
	reqValidator := form_requests.NewBookUpdateRequest().Validate(&req, ctx.Context())
	if reqValidator != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"errors": reqValidator, "message": "Validation failed", "status": http.StatusBadRequest})
	}

	params := &dtos.GetBookParams{ID: req.ID}
	// Fetch the existing book to get the current data
	existingBook, err := c.service.GetBookByID(ctx.Context(), params)
	if err != nil {
		return utils.GetResponse(ctx, nil, nil, "Book not found", http.StatusNotFound, err.Error(), nil)
	}

	// Extract user ID from JWT
	claims, err := middleware.GetAuthUser(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"errors": err.Error(), "message": "Unbookized", "status": fiber.StatusUnauthorized})
	}
	userID := uint(claims["user_id"].(float64))

	book := models.Book{
		ID:          req.ID,
		Name:        req.Name,
		Description: req.Description,
		AuthorID:    req.AuthorID,
		CategoryID:  req.CategoryID,
		CreatedByID: existingBook.CreatedByID,
		UpdatedByID: &userID,
		CreatedAt:   existingBook.CreatedAt,
	}

	updatedBook, err := c.service.UpdateBook(ctx.Context(), &book)
	if err != nil {
		return utils.GetResponse(ctx, nil, nil, "Failed to update book", http.StatusInternalServerError, err.Error(), nil)
	}

	params = &dtos.GetBookParams{ID: updatedBook.ID}
	getBook, err := c.service.GetBookByID(ctx.Context(), params)
	if err != nil {
		return utils.GetResponse(ctx, nil, nil, "Book not found", http.StatusNotFound, err.Error(), nil)
	}

	filters := ctx.Locals("filters").(map[string]string)
	paginationMeta := utils.CreatePaginationMeta(filters, 1)

	return utils.GetResponse(ctx, []interface{}{getBook}, paginationMeta, "Book updated successfully", http.StatusOK, nil, nil)
}

// delete book
func (c *BookController) DeleteBook(ctx *fiber.Ctx) error {
	var req dtos.DeleteBookRequest

	if err := ctx.BodyParser(&req); err != nil {
		return utils.GetResponse(ctx, nil, nil, "Book not found", http.StatusBadRequest, err.Error(), nil)
	}

	if req.ID == 0 {
		return utils.GetResponse(ctx, nil, nil, "Book not found", http.StatusBadRequest, "ID is required", nil)
	}

	params := &dtos.GetBookParams{ID: req.ID}
	// GET book by ID
	_, err := c.service.GetBookByID(ctx.Context(), params)
	if err != nil {
		return utils.GetResponse(ctx, nil, nil, "Book not found", http.StatusNotFound, err.Error(), nil)
	}

	err = c.service.DeleteBook(ctx.Context(), req.ID)
	if err != nil {
		return utils.GetResponse(ctx, nil, nil, "Failed to delete book", http.StatusInternalServerError, err.Error(), nil)
	}

	return utils.GetResponse(ctx, nil, nil, "Book deleted successfully", http.StatusOK, nil, nil)
}

// restore book
func (c *BookController) RestoreBook(ctx *fiber.Ctx) error {
	var req dtos.DeleteBookRequest

	if err := ctx.BodyParser(&req); err != nil {
		return utils.GetResponse(ctx, nil, nil, "Book not found", http.StatusBadRequest, err.Error(), nil)
	}

	if req.ID == 0 {
		return utils.GetResponse(ctx, nil, nil, "Book not found", http.StatusBadRequest, "ID is required", nil)
	}

	isDeleted := 1
	params := &dtos.GetBookParams{ID: req.ID, IsDeleted: &isDeleted}
	// GET book by ID
	_, err := c.service.GetBookByID(ctx.Context(), params)
	if err != nil {
		return utils.GetResponse(ctx, nil, nil, "Book not found", http.StatusNotFound, err.Error(), nil)
	}

	err = c.service.RestoreBook(ctx.Context(), req.ID)
	if err != nil {
		return utils.GetResponse(ctx, nil, nil, "Failed to restore book", http.StatusInternalServerError, err.Error(), nil)
	}

	return utils.GetResponse(ctx, nil, nil, "Book restored successfully", http.StatusOK, nil, nil)
}
