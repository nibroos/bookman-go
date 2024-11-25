package service

import (
	"context"

	"github.com/nibroos/bookman-go/user-service/internal/dtos"
	"github.com/nibroos/bookman-go/user-service/internal/models"
	"github.com/nibroos/bookman-go/user-service/internal/repository"
)

type AuthorService struct {
	repo *repository.AuthorRepository
}

func NewAuthorService(repo *repository.AuthorRepository) *AuthorService {
	return &AuthorService{repo: repo}
}

func (s *AuthorService) GetAuthors(ctx context.Context, filters map[string]string) ([]dtos.AuthorListDTO, int, error) {

	resultChan := make(chan dtos.GetAuthorsResult, 1)

	go func() {
		authors, total, err := s.repo.GetAuthors(ctx, filters)
		resultChan <- dtos.GetAuthorsResult{Authors: authors, Total: total, Err: err}
	}()

	select {
	case res := <-resultChan:
		return res.Authors, res.Total, res.Err
	case <-ctx.Done():
		return nil, 0, ctx.Err()
	}
}

func (s *AuthorService) CreateAuthor(ctx context.Context, author *models.Author) (*models.Author, error) {
	// Transaction handling
	tx := s.repo.BeginTransaction()
	if err := tx.Error; err != nil {
		return nil, err
	}

	// Create author
	if err := s.repo.CreateAuthor(tx, author); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return author, nil
}

func (s *AuthorService) GetAuthorByID(ctx context.Context, params *dtos.GetAuthorParams) (*dtos.AuthorDetailDTO, error) {
	authorChan := make(chan *dtos.AuthorDetailDTO, 1)
	errChan := make(chan error, 1)

	go func() {
		author, err := s.repo.GetAuthorByID(ctx, params)
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

func (s *AuthorService) UpdateAuthor(ctx context.Context, author *models.Author) (*models.Author, error) {
	// Transaction handling
	tx := s.repo.BeginTransaction()
	if err := tx.Error; err != nil {
		return nil, err
	}

	// Update author
	if err := s.repo.UpdateAuthor(tx, author); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return author, nil
}

func (s *AuthorService) DeleteAuthor(ctx context.Context, id uint) error {
	// Transaction handling
	tx := s.repo.BeginTransaction()
	if err := tx.Error; err != nil {
		return err
	}

	// Delete author
	if err := s.repo.DeleteAuthor(tx, id); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func (s *AuthorService) RestoreAuthor(ctx context.Context, id uint) error {
	// Transaction handling
	tx := s.repo.BeginTransaction()
	if err := tx.Error; err != nil {
		return err
	}

	// Restore author
	if err := s.repo.RestoreAuthor(tx, id); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
