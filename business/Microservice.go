package business

import (
	"gorm.io/gorm"
)

// Microservice struct represents a microservice with its properties.
type Microservice struct {
	gorm.Model   // This includes fields ID, CreatedAt, UpdatedAt, DeletedAt
	FriendlyName string
	RepoLink     string
	Status       string
	UserID       uint
	Inputs       []Input
	OutputLink   string
	BackendName  string `gorm:"unique"`
	ImageID      string
}
