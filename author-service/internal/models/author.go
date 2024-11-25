package models

import (
	"time"

	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	ID          uint       `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Name        string     `json:"name" gorm:"column:name"`
	Description string     `json:"desc" gorm:"column:description"`
	CreatedByID *uint      `json:"created_by_id" gorm:"column:created_by_id"`
	UpdatedByID *uint      `json:"updated_by_id" gorm:"column:updated_by_id"`
	CreatedAt   *time.Time `json:"created_at" gorm:"column:created_at"`
	DeletedAt   *time.Time `json:"deleted_at" gorm:"column:deleted_at"`
	Version     uint       `json:"version" gorm:"column:version"`
}
