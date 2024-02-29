package dataaccess

import (
	"github.com/GoKubes/ServerlessOrchestrator/business"
	"gorm.io/gorm"
)

type MicroservicesDAOpq struct {
	db *gorm.DB
}

// NewMicroservicesDAO creates a new instance of MicroservicesDAO with a database connection.
func NewMicroservicesDAO(db *gorm.DB) *MicroservicesDAOpq {
	return &MicroservicesDAOpq{db: db}
}

// GetAll retrieves all microservices from the database.
func (dao *MicroservicesDAOpq) GetAll() ([]business.Microservice, error) {
	var microservices []business.Microservice
	result := dao.db.Find(&microservices)
	if result.Error != nil {
		return nil, result.Error
	}
	return microservices, nil
}

// Insert inserts a microservice and its associated inputs into the database transactionally
func (dao *MicroservicesDAOpq) Insert(micro business.Microservice) error {
	// Start a transaction
	tx := dao.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// Insert the microservice into the database
	if err := tx.Create(&micro).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Insert the inputs associated with the microservice into the database
	for _, input := range micro.Inputs {
		input.MicroserviceID = micro.ID // Assign the MicroserviceID
		if err := tx.Create(&input).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

// Delete removes a microservice record from the database by ID.
func (dao *MicroservicesDAOpq) Delete(id uint) error {
	result := dao.db.Delete(&business.Microservice{}, id)
	return result.Error
}

// GetByID retrieves a microservice by its ID.
func (dao *MicroservicesDAOpq) GetByID(id uint) (business.Microservice, error) {
	var micro business.Microservice
	result := dao.db.First(&micro, id)
	if result.Error != nil {
		return business.Microservice{}, result.Error
	}
	return micro, nil
}

// GetByName retrieves a microservice by its name, case-insensitively.
func (dao *MicroservicesDAOpq) GetByName(name string) (business.Microservice, error) {
	var micro business.Microservice
	// Using ILIKE for PostgreSQL for case-insensitive comparison
	result := dao.db.Where("LOWER(name) = LOWER(?)", name).First(&micro)
	if result.Error != nil {
		return business.Microservice{}, result.Error
	}
	return micro, nil
}

// Update modifies an existing microservice record.
func (dao *MicroservicesDAOpq) Update(micro business.Microservice) error {
	result := dao.db.Save(&micro)
	return result.Error
}

// Ensure that MicroservicesDAOpq implements the MicroservicesDAO_IF interface.
var _ business.DAO_IF = &MicroservicesDAOpq{}
