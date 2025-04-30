package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User represents a user in the system
type User struct {
	ID       string `json:"id" gorm:"type:char(36);primary_key"` // Gunakan CHAR(36) untuk UUID
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// BeforeCreate is a GORM hook to generate UUID before saving data
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	// Generate UUID before saving as string
	u.ID = uuid.New().String()
	return nil
}
