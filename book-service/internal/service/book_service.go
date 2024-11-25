package service

import (
	"context"

	"github.com/nibroos/bookman-go/user-service/internal/dtos"
	"github.com/nibroos/bookman-go/user-service/internal/models"
	"github.com/nibroos/bookman-go/user-service/internal/repository"
)

type BookService struct {
	repo *repository.BookRepository
}

func NewBookService(repo *repository.BookRepository) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) GetBooks(ctx context.Context, filters map[string]string) ([]dtos.BookListDTO, int, error) {

	resultChan := make(chan dtos.GetBooksResult, 1)

	go func() {
		books, total, err := s.repo.GetBooks(ctx, filters)
		resultChan <- dtos.GetBooksResult{Books: books, Total: total, Err: err}
	}()

	select {
	case res := <-resultChan:
		return res.Books, res.Total, res.Err
	case <-ctx.Done():
		return nil, 0, ctx.Err()
	}
}

func (s *BookService) CreateBook(ctx context.Context, book *models.Book) (*models.Book, error) {
	// Transaction handling
	tx := s.repo.BeginTransaction()
	if err := tx.Error; err != nil {
		return nil, err
	}

	// Create book
	if err := s.repo.CreateBook(tx, book); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return book, nil
}

func (s *BookService) GetBookByID(ctx context.Context, params *dtos.GetBookParams) (*dtos.BookDetailDTO, error) {
	bookChan := make(chan *dtos.BookDetailDTO, 1)
	errChan := make(chan error, 1)

	go func() {
		book, err := s.repo.GetBookByID(ctx, params)
		if err != nil {
			errChan <- err
			return
		}
		bookChan <- book
	}()

	select {
	case book := <-bookChan:
		return book, nil
	case err := <-errChan:
		return nil, err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func (s *BookService) UpdateBook(ctx context.Context, book *models.Book) (*models.Book, error) {
	// Transaction handling
	tx := s.repo.BeginTransaction()
	if err := tx.Error; err != nil {
		return nil, err
	}

	// Update book
	if err := s.repo.UpdateBook(tx, book); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return book, nil
}

func (s *BookService) DeleteBook(ctx context.Context, id uint) error {
	// Transaction handling
	tx := s.repo.BeginTransaction()
	if err := tx.Error; err != nil {
		return err
	}

	// Delete book
	if err := s.repo.DeleteBook(tx, id); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func (s *BookService) RestoreBook(ctx context.Context, id uint) error {
	// Transaction handling
	tx := s.repo.BeginTransaction()
	if err := tx.Error; err != nil {
		return err
	}

	// Restore book
	if err := s.repo.RestoreBook(tx, id); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
