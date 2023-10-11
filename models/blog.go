package models

import (
	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model
	Title        string `json:"title" gorm:"unique"`
	CreatedBy    int
	Body         string `json:"body"`
	User         User   `gorm:"foreignKey:CreatedBy" json:"-"`
	CategoryId   int    `json:"category_id"`
	CategoryName string `json:"category_name"`

	// // Meta
	// CreatedAt time.Time `json:"-"`
	// UpdatedAt time.Time `json:"-"`
}
