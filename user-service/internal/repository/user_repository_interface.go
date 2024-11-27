package repository

import (
	"context"

	"github.com/nibroos/bookman-go/user-service/internal/dtos"
)

type UserRepositoryInterface interface {
    GetUsers(ctx context.Context, filters map[string]string) ([]dtos.UserListDTO, int, error)
}