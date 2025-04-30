package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Todo represents a todo item
type Todo struct {
	ID     string `json:"id" gorm:"type:char(36);primary_key"` // Gunakan CHAR(36) untuk UUID
	Title  string `json:"title"`
	Status string `json:"status"`
}

// BeforeCreate is a GORM hook to generate UUID before saving data
func (t *Todo) BeforeCreate(tx *gorm.DB) (err error) {
	// Generate UUID sebelum menyimpan data
	t.ID = uuid.New().String()
	return nil
}
