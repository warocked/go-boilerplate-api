package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User represents a user in the database
type User struct {
	ID        string    `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Email     string    `json:"email" gorm:"type:varchar(255);uniqueIndex;not null"`
	FirstName string    `json:"first_name" gorm:"type:varchar(100)"`
	LastName  string    `json:"last_name" gorm:"type:varchar(100)"`
	PasswordHash string `json:"-" gorm:"type:varchar(255);not null;column:password_hash"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// TableName specifies the table name for GORM
func (User) TableName() string {
	return "users"
}

// BeforeCreate hook to generate UUID if not set
func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.ID == "" {
		u.ID = uuid.New().String()
	}
	return nil
}
