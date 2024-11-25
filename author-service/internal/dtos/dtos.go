package dtos

import (
	"time"
)

type CreateAuthorRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	// AttachmentUrls []string `json:"attachment_urls"`
}

type UpdateAuthorRequest struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	// AttachmentUrls []string `json:"attachment_urls"`
}

type GetAuthorByIDRequest struct {
	ID uint `json:"id"`
}

type GetAuthorParams struct {
	ID        uint
	IsDeleted *int
}

func NewGetAuthorParams(id uint) *GetAuthorParams {
	defaultIsDeleted := 0
	return &GetAuthorParams{
		ID:        id,
		IsDeleted: &defaultIsDeleted,
	}
}

type DeleteAuthorRequest struct {
	ID uint `json:"id"`
}

type AuthorListDTO struct {
	ID            int     `json:"id" db:"id"`
	Name          string  `json:"name" db:"name"`
	Description   string  `json:"description" db:"description"`
	CreatedByName *string `json:"created_by_name" db:"created_by_name"`
	UpdatedByName *string `json:"updated_by_name" db:"updated_by_name"`
	CreatedAt     *string `json:"created_at" db:"created_at"`
	UpdatedAt     *string `json:"updated_at" db:"updated_at"`
	DeleteAt      *string `json:"deleted_at" db:"deleted_at"`
}

type AuthorDetailDTO struct {
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
type GetAuthorsResult struct {
	Authors []AuthorListDTO
	Total   int
	Err     error
}
