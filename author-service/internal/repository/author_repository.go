package repository

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/nibroos/bookman-go/user-service/internal/dtos"
	"github.com/nibroos/bookman-go/user-service/internal/models"
	"github.com/nibroos/bookman-go/user-service/internal/utils"
	"gorm.io/gorm"
)

type AuthorRepository struct {
	db    *gorm.DB
	sqlDB *sqlx.DB
}

func NewAuthorRepository(db *gorm.DB, sqlDB *sqlx.DB) *AuthorRepository {
	return &AuthorRepository{
		db:    db,
		sqlDB: sqlDB,
	}
}

func (r *AuthorRepository) GetAuthors(ctx context.Context, filters map[string]string) ([]dtos.AuthorListDTO, int, error) {
	authors := []dtos.AuthorListDTO{}
	var total int

	query := `SELECT *
    FROM ( 
        SELECT e.id, e.name, e.description, e.created_at, e.updated_at, e.deleted_at

        FROM authors e
    ) AS alias WHERE 1=1 AND deleted_at IS NULL`

	countQuery := `SELECT COUNT(*) FROM (
        SELECT e.id, e.name, e.description, e.created_at, e.updated_at, e.deleted_at

        FROM authors e
    ) AS alias WHERE 1=1 AND deleted_at IS NULL`

	var args []interface{}

	i := 1
	for key, value := range filters {
		switch key {
		case "name", "description":
			if value != "" {
				query += fmt.Sprintf(" AND %s ILIKE $%d", key, i)
				countQuery += fmt.Sprintf(" AND %s ILIKE $%d", key, i)
				args = append(args, "%"+value+"%")
				i++
			}
		}
	}

	if value, ok := filters["global"]; ok && value != "" {
		query += fmt.Sprintf(" AND (name ILIKE $%d OR description ILIKE $%d)", i, i+1)
		countQuery += fmt.Sprintf(" AND (name ILIKE $%d OR description ILIKE $%d)", i, i+1)
		args = append(args, "%"+value+"%", "%"+value+"%")
		i += 2
	}

	countArgs := append([]interface{}{}, args...)

	// Channels for concurrent execution
	countChan := make(chan error)
	selectChan := make(chan error)

	// Goroutine for count query
	go func() {
		err := r.sqlDB.GetContext(ctx, &total, countQuery, countArgs...)
		countChan <- err
	}()

	orderColumn := utils.GetStringOrDefault(filters["order_column"], "id")
	orderDirection := utils.GetStringOrDefault(filters["order_direction"], "asc")
	query += fmt.Sprintf(" ORDER BY %s %s", orderColumn, orderDirection)

	perPage := utils.GetIntOrDefault(filters["per_page"], 10)
	currentPage := utils.GetIntOrDefault(filters["page"], 1)

	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", i, i+1)
	args = append(args, perPage, (currentPage-1)*perPage)

	// Goroutine for select query
	go func() {
		err := r.sqlDB.SelectContext(ctx, &authors, query, args...)
		selectChan <- err
	}()

	// Wait for both goroutines to finish
	countErr := <-countChan
	selectErr := <-selectChan

	if countErr != nil {
		return nil, 0, countErr
	}

	if selectErr != nil {
		return nil, 0, selectErr
	}

	return authors, total, nil
}

func (r *AuthorRepository) GetAuthorByID(ctx context.Context, params *dtos.GetAuthorParams) (*dtos.AuthorDetailDTO, error) {
	var author dtos.AuthorDetailDTO
	// deletedAt := params.IsDeleted

	query := `SELECT e.*

	FROM authors e
	WHERE 1=1`

	var args []interface{}

	i := 1
	query += " AND e.id = $1"
	args = append(args, params.ID)
	i++

	isDeletedQuery := ` AND e.deleted_at IS NULL`
	if params.IsDeleted != nil && *params.IsDeleted == 1 {
		isDeletedQuery = " AND e.deleted_at IS NOT NULL"
	}

	query += isDeletedQuery

	if err := r.sqlDB.Get(&author, query, args...); err != nil {
		return nil, err
	}

	return &author, nil
}

// BeginTransaction starts a new transaction
func (r *AuthorRepository) BeginTransaction() *gorm.DB {
	return r.db.Begin()
}

func (r *AuthorRepository) CreateAuthor(tx *gorm.DB, author *models.Author) error {
	if err := tx.Create(author).Error; err != nil {
		return err
	}
	return nil
}

func (r *AuthorRepository) UpdateAuthor(tx *gorm.DB, author *models.Author) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(author).Error; err != nil {
			return err
		}
		return nil
	})

}

func (r *AuthorRepository) DeleteAuthor(tx *gorm.DB, id uint) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// if err := tx.Unscoped().Delete(&models.Author{}, id).Error; err != nil {
		if err := tx.Delete(&models.Author{}, id).Error; err != nil {
			return err
		}
		return nil
	})
}

func (r *AuthorRepository) RestoreAuthor(tx *gorm.DB, id uint) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("UPDATE authors SET deleted_at = NULL WHERE id = ?", id).Error; err != nil {
			return err
		}
		return nil
	})
}
