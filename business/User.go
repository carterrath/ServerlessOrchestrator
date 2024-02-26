package business

import (
	"gorm.io/gorm"
)

// User struct represents a developer or consumer with its properties.
type User struct {
	gorm.Model        // This includes fields ID, CreatedAt, UpdatedAt, DeletedAt
	Username   string `gorm:"unique"`
	Email      string `gorm:"unique"`
	Password   string
	UserType   string // Developer or Consumer
}
