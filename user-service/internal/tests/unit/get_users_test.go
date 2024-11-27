package unit_test

import (
	"context"
	"testing"
	"time"

	"github.com/nibroos/bookman-go/user-service/internal/dtos"
	"github.com/nibroos/bookman-go/user-service/internal/models"
	"github.com/nibroos/bookman-go/user-service/internal/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

// MockUserRepository is a mock implementation of the UserRepository interface
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) GetUsers(ctx context.Context, filters map[string]string) ([]dtos.UserListDTO, int, error) {
	args := m.Called(ctx, filters)
	return args.Get(0).([]dtos.UserListDTO), args.Int(1), args.Error(2)
}

func (m *MockUserRepository) GetUserByID(ctx context.Context, params *dtos.GetUserByIDParams) (*dtos.UserDetailDTO, error) {
    args := m.Called(ctx, params)
    return args.Get(0).(*dtos.UserDetailDTO), args.Error(1)
}

func (m *MockUserRepository) GetUserByEmail(ctx context.Context, email string) (*dtos.UserDetailDTO, error) {
    args := m.Called(ctx, email)
    return args.Get(0).(*dtos.UserDetailDTO), args.Error(1)
}

func (m *MockUserRepository) BeginTransaction() *gorm.DB {
    args := m.Called()
    return args.Get(0).(*gorm.DB)
}

func (m *MockUserRepository) AttachRoles(tx *gorm.DB, user *models.User, roleIDs []uint32) error {
    args := m.Called(tx, user, roleIDs)
    return args.Error(0)
}

func (m *MockUserRepository) CreateUser(tx *gorm.DB, user *models.User) error {
    args := m.Called(tx, user)
    return args.Error(0)
}

func (m *MockUserRepository) UpdateUser(tx *gorm.DB, user *models.User) error {
    args := m.Called(tx, user)
    return args.Error(0)
}

func (m *MockUserRepository) DeleteUser(tx *gorm.DB, id uint) error {
    args := m.Called(tx, id)
    return args.Error(0)
}

func (m *MockUserRepository) DeleteRolesByUserID(tx *gorm.DB, userID uint) error {
    args := m.Called(tx, userID)
    return args.Error(0)
}

func (m *MockUserRepository) RestoreUser(tx *gorm.DB, id uint) error {
    args := m.Called(tx, id)
    return args.Error(0)
}

func TestGetUsers(t *testing.T) {
    mockRepo := new(MockUserRepository)
    userService := service.NewUserService(mockRepo)

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    ptr := func(s string) *string { return &s }
    users := []dtos.UserListDTO{
        {ID: 1, Username: ptr("user1"), Name: "User One", Email: "user1@example.com"},
        {ID: 2, Username: ptr("user2"), Name: "User Two", Email: "user2@example.com"},
    }

    tests := []struct {
        name     string
        filters  map[string]string
        mockResp []dtos.UserListDTO
        mockTotal int
        mockErr  error
        expectedUsers []dtos.UserListDTO
        expectedTotal int
        expectedErr   error
    }{
        {
            name: "returned data without filter",
            filters: map[string]string{
                "page":            "1",
                "global":          "",
                "name":            "",
                "per_page":        "10",
                "order_column":    "",
                "order_direction": "desc",
            },
            mockResp: users,
            mockTotal: 2,
            mockErr:  nil,
            expectedUsers: users,
            expectedTotal: 2,
            expectedErr:   nil,
        },
        {
            name: "returned data with filter",
            filters: map[string]string{
                "page":            "1",
                "global":          "",
                "name":            "User One",
                "per_page":        "10",
                "order_column":    "",
                "order_direction": "desc",
            },
            mockResp: []dtos.UserListDTO{users[0]},
            mockTotal: 1,
            mockErr:  nil,
            expectedUsers: []dtos.UserListDTO{users[0]},
            expectedTotal: 1,
            expectedErr:   nil,
        },
        {
            name: "returned no data with filter that total 0",
            filters: map[string]string{
                "page":            "1",
                "global":          "",
                "name":            "Nonexistent User",
                "per_page":        "10",
                "order_column":    "",
                "order_direction": "desc",
            },
            mockResp: []dtos.UserListDTO{},
            mockTotal: 0,
            mockErr:  nil,
            expectedUsers: []dtos.UserListDTO{},
            expectedTotal: 0,
            expectedErr:   nil,
        },
        {
            name: "repository returns error",
            filters: map[string]string{
                "page":            "1",
                "global":          "",
                "name":            "",
                "per_page":        "10",
                "order_column":    "",
                "order_direction": "desc",
            },
            mockResp: nil,
            mockTotal: 0,
            mockErr:  assert.AnError,
            expectedUsers: nil,
            expectedTotal: 0,
            expectedErr:   assert.AnError,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            mockRepo.On("GetUsers", ctx, tt.filters).Return(tt.mockResp, tt.mockTotal, tt.mockErr).Once()
            resultUsers, total, err := userService.GetUsers(ctx, tt.filters)
            assert.Equal(t, tt.expectedErr, err)
            assert.Equal(t, tt.expectedTotal, total)
            assert.Equal(t, tt.expectedUsers, resultUsers)
            mockRepo.AssertExpectations(t)
        })
    }
}