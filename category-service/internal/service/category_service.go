package service

import (
	"context"

	"github.com/nibroos/bookman-go/user-service/internal/dtos"
	"github.com/nibroos/bookman-go/user-service/internal/models"
	"github.com/nibroos/bookman-go/user-service/internal/repository"
)

type CategoryService struct {
	repo *repository.CategoryRepository
}

func NewCategoryService(repo *repository.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) GetCategories(ctx context.Context, filters map[string]string) ([]dtos.CategoryListDTO, int, error) {

	resultChan := make(chan dtos.GetCategoriesResult, 1)

	go func() {
		categories, total, err := s.repo.GetCategories(ctx, filters)
		resultChan <- dtos.GetCategoriesResult{Categories: categories, Total: total, Err: err}
	}()

	select {
	case res := <-resultChan:
		return res.Categories, res.Total, res.Err
	case <-ctx.Done():
		return nil, 0, ctx.Err()
	}
}

func (s *CategoryService) CreateCategory(ctx context.Context, author *models.Category) (*models.Category, error) {
	// Transaction handling
	tx := s.repo.BeginTransaction()
	if err := tx.Error; err != nil {
		return nil, err
	}

	// Create author
	if err := s.repo.CreateCategory(tx, author); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return author, nil
}

func (s *CategoryService) GetCategoryByID(ctx context.Context, params *dtos.GetCategoryParams) (*dtos.CategoryDetailDTO, error) {
	authorChan := make(chan *dtos.CategoryDetailDTO, 1)
	errChan := make(chan error, 1)

	go func() {
		author, err := s.repo.GetCategoryByID(ctx, params)
		if err != nil {
			errChan <- err
			return
		}
		authorChan <- author
	}()

	select {
	case author := <-authorChan:
		return author, nil
	case err := <-errChan:
		return nil, err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func (s *CategoryService) UpdateCategory(ctx context.Context, author *models.Category) (*models.Category, error) {
	// Transaction handling
	tx := s.repo.BeginTransaction()
	if err := tx.Error; err != nil {
		return nil, err
	}

	// Update author
	if err := s.repo.UpdateCategory(tx, author); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return author, nil
}

func (s *CategoryService) DeleteCategory(ctx context.Context, id uint) error {
	// Transaction handling
	tx := s.repo.BeginTransaction()
	if err := tx.Error; err != nil {
		return err
	}

	// Delete author
	if err := s.repo.DeleteCategory(tx, id); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func (s *CategoryService) RestoreCategory(ctx context.Context, id uint) error {
	// Transaction handling
	tx := s.repo.BeginTransaction()
	if err := tx.Error; err != nil {
		return err
	}

	// Restore author
	if err := s.repo.RestoreCategory(tx, id); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
