package dtos

import (
	"time"
)

type CreateBookRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	CategoryID  uint   `json:"category_id"`
	AuthorID    uint   `json:"author_id"`
}

type UpdateBookRequest struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CategoryID  uint   `json:"category_id"`
	AuthorID    uint   `json:"author_id"`
}

type GetBookByIDRequest struct {
	ID uint `json:"id"`
}

type GetBookParams struct {
	ID        uint
	IsDeleted *int
}

func NewGetBookParams(id uint) *GetBookParams {
	defaultIsDeleted := 0
	return &GetBookParams{
		ID:        id,
		IsDeleted: &defaultIsDeleted,
	}
}

type DeleteBookRequest struct {
	ID uint `json:"id"`
}

type BookListDTO struct {
	ID            int     `json:"id" db:"id"`
	Name          string  `json:"name" db:"name"`
	Description   string  `json:"description" db:"description"`
	CategoryID    uint    `json:"category_id" db:"category_id"`
	AuthorID      uint    `json:"author_id" db:"author_id"`
	CreatedByName *string `json:"created_by_name" db:"created_by_name"`
	UpdatedByName *string `json:"updated_by_name" db:"updated_by_name"`
	CreatedAt     *string `json:"created_at" db:"created_at"`
	UpdatedAt     *string `json:"updated_at" db:"updated_at"`
	DeleteAt      *string `json:"deleted_at" db:"deleted_at"`
}

type BookDetailDTO struct {
	ID            uint       `json:"id" db:"id"`
	Name          string     `json:"name" db:"name"`
	Description   string     `json:"description" db:"description"`
	CategoryID    uint       `json:"category_id" db:"category_id"`
	AuthorID      uint       `json:"author_id" db:"author_id"`
	CreatedByID   uint       `json:"created_by_id" db:"created_by_id"`
	UpdatedByID   *uint      `json:"updated_by_id" db:"updated_by_id"`
	CreatedByName *string    `json:"created_by_name" db:"created_by_name"`
	UpdatedByName *string    `json:"updated_by_name" db:"updated_by_name"`
	CreatedAt     *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at" db:"deleted_at"`
}
type GetBooksResult struct {
	Books []BookListDTO
	Total int
	Err   error
}
