package dtos

import (
	"time"
)

type CreateCategoryRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	// AttachmentUrls []string `json:"attachment_urls"`
}

type UpdateCategoryRequest struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	// AttachmentUrls []string `json:"attachment_urls"`
}

type GetCategoryByIDRequest struct {
	ID uint `json:"id"`
}

type GetCategoryParams struct {
	ID        uint
	IsDeleted *int
}

func NewGetCategoryParams(id uint) *GetCategoryParams {
	defaultIsDeleted := 0
	return &GetCategoryParams{
		ID:        id,
		IsDeleted: &defaultIsDeleted,
	}
}

type DeleteCategoryRequest struct {
	ID uint `json:"id"`
}

type CategoryListDTO struct {
	ID            int     `json:"id" db:"id"`
	Name          string  `json:"name" db:"name"`
	Description   string  `json:"description" db:"description"`
	CreatedByName *string `json:"created_by_name" db:"created_by_name"`
	UpdatedByName *string `json:"updated_by_name" db:"updated_by_name"`
	CreatedAt     *string `json:"created_at" db:"created_at"`
	UpdatedAt     *string `json:"updated_at" db:"updated_at"`
	DeleteAt      *string `json:"deleted_at" db:"deleted_at"`
}

type CategoryDetailDTO struct {
	ID            uint       `json:"id" db:"id"`
	Name          string     `json:"name" db:"name"`
	Description   string     `json:"description" db:"description"`
	CreatedByID   uint       `json:"created_by_id" db:"created_by_id"`
	UpdatedByID   *uint      `json:"updated_by_id" db:"updated_by_id"`
	CreatedByName *string    `json:"created_by_name" db:"created_by_name"`
	UpdatedByName *string    `json:"updated_by_name" db:"updated_by_name"`
	CreatedAt     *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at" db:"deleted_at"`
}
type GetCategoriesResult struct {
	Categories []CategoryListDTO
	Total      int
	Err        error
}
