package business

import (
	"gorm.io/gorm"
)

// Microservice struct represents a microservice with its properties.
type Microservice struct {
	gorm.Model         // This includes fields ID, CreatedAt, UpdatedAt, DeletedAt
	Name        string `gorm:"unique"`
	PlaceHolder string
	Status      string
	Author      string
	Inputs      []Input
}
