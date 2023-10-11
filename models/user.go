package models

import (
	"gorm.io/gorm"
)

type UserRole struct {
	Id   int    `json:"id"`
	Role string `json:"role"`
}

type User struct {
	gorm.Model
	ID         int    `json:"id"`
	FirstName  string `json:"firstname"`
	LastName   string `json:"lastname"`
	Email      string `json:"email" gorm:"unique"`
	UserRoleID int    `json:"role_id"`
	Password   string `json:"-"`

	// CreatedAt time.Time `json:"created_at"`
	// UpdatedAt time.Time `json:"updated_at,timestamp"`
	IsActive bool `json:"is_active,boolean"`
}
