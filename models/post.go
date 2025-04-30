// models/post.go

package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Post represents a post in the system
type Post struct {
	ID          string `json:"id" gorm:"type:char(36);primary_key"` // Gunakan CHAR(36) untuk UUID
	Title       string `json:"title"`
	Image       string `json:"image"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

// BeforeCreate is a GORM hook to generate UUID before saving data
func (p *Post) BeforeCreate(tx *gorm.DB) (err error) {
	// Generate UUID before saving as string
	p.ID = uuid.New().String()
	return nil
}
