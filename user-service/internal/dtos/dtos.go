package dtos

import (
	"time"

	"github.com/nibroos/bookman-go/user-service/internal/utils"
)

type GetUsersRequest struct {
	Global         string                 `json:"global"`
	Username       string                 `json:"username"`
	Name           string                 `json:"name"`
	Email          string                 `json:"email"`
	PerPage        *string `json:"per_page" default:"10"`         // Default per_page to 10
	Page           *string `json:"page" default:"1"`              // Default page to 1
	OrderColumn    string                 `json:"order_column" default:"id"`     // Default order column to "id"
	OrderDirection string                 `json:"order_direction" default:"asc"` // Default order direction to "asc"
}

type CreateUserRequest struct {
	Name     string                 `json:"name"`
	Username *string `json:"username"`
	Email    string                 `json:"email"`
	Address  *string `json:"address"`
	Password string                 `json:"password"`
	RoleIDs  []uint32               `json:"role_ids"`
}

type UpdateUserRequest struct {
	ID       uint                   `json:"id"`
	Username *string `json:"username"`
	Name     string                 `json:"name"`
	Email    string                 `json:"email"`
	Address  *string `json:"address"`
	Password *string `json:"password"`
	RoleIDs  []uint32               `json:"role_ids"`
}

type GetUserByIDParams struct {
	ID        uint `json:"id"`
	IsDeleted *int
}

type GetUserByIDRequest struct {
	ID uint `json:"id"`
}

type DeleteUserRequest struct {
	ID uint `json:"id"`
}

type UserListDTO struct {
	ID       int                    `json:"id"`
	Username *string `json:"username"`
	Name     string                 `json:"name"`
	Email    string                 `json:"email"`
}

type UserDetailDTO struct {
	ID          uint                   `json:"id"`
	Name        string                 `json:"name"`
	Username    *string `json:"username"`
	Email       string                 `json:"email"`
	Address     *string `json:"address"`
	Password    *string                `json:"password"`
	Roles       []string               `json:"roles"`
	Permissions []string               `json:"permissions"`
	CreatedAt   *string                `json:"created_at"`
}
type GetUsersResult struct {
	Users []UserListDTO
	Total int
	Err   error
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type RegisterRequest struct {
	Name     string                 `json:"name"`
	Username *string `json:"username"`
	Email    string                 `json:"email"`
	Address  *string `json:"address"`
	Password string                 `json:"password"`
}

type CreateAuthorRequest struct {
	ModuleID    uint   `json:"module_id"`
	NoUrut      uint   `json:"no_urut"`
	Name        string `json:"name"`
	Description string `json:"description"`
	TextMateri  string `json:"text_materi"`
	// AttachmentUrls []string `json:"attachment_urls"`
}

type UpdateAuthorRequest struct {
	ID          uint   `json:"id"`
	ModuleID    uint   `json:"module_id"`
	NoUrut      uint   `json:"no_urut"`
	Name        string `json:"name"`
	Description string `json:"description"`
	TextMateri  string `json:"text_materi"`
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
	NoUrut        uint    `json:"no_urut" db:"no_urut"`
	Name          string  `json:"name" db:"name"`
	Description   string  `json:"description" db:"description"`
	ThumbnailURL  *string `json:"thumbnail_url" db:"thumbnail_url"`
	VideoURL      *string `json:"video_url" db:"video_url"`
	ModuleID      uint    `json:"module_id" db:"module_id"`
	ModuleName    string  `json:"module_name" db:"module_name"`
	TextMaterial  string  `json:"text_materi" db:"text_materi"`
	CreatedByName *string `json:"created_by_name" db:"created_by_name"`
	UpdatedByName *string `json:"updated_by_name" db:"updated_by_name"`
	CreatedAt     *string `json:"created_at" db:"created_at"`
	UpdatedAt     *string `json:"updated_at" db:"updated_at"`
	DeleteAt      *string `json:"deleted_at" db:"deleted_at"`
}

type AuthorDetailDTO struct {
	ID             uint                  `json:"id" db:"id"`
	NoUrut         uint                  `json:"no_urut" db:"no_urut"`
	Name           string                `json:"name" db:"name"`
	Description    string                `json:"description" db:"description"`
	ThumbnailURL   *string               `json:"thumbnail_url" db:"thumbnail_url"`
	VideoURL       *string               `json:"video_url" db:"video_url"`
	AttachmentURLs utils.JSONStringArray `json:"attachment_urls" db:"attachment_urls"`
	ModuleID       uint                  `json:"module_id" db:"module_id"`
	ModuleName     string                `json:"module_name" db:"module_name"`
	TextMaterial   string                `json:"text_materi" db:"text_materi"`
	CreatedByID    uint                  `json:"created_by_id" db:"created_by_id"`
	UpdatedByID    *uint                 `json:"updated_by_id" db:"updated_by_id"`
	CreatedByName  *string               `json:"created_by_name" db:"created_by_name"`
	UpdatedByName  *string               `json:"updated_by_name" db:"updated_by_name"`
	CreatedAt      *time.Time            `json:"created_at" db:"created_at"`
	UpdatedAt      *time.Time            `json:"updated_at" db:"updated_at"`
	DeletedAt      *time.Time            `json:"deleted_at" db:"deleted_at"`
}
type GetAuthorsResult struct {
	Authors []AuthorListDTO
	Total   int
	Err     error
}
