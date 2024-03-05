package dataaccess

import (
	"errors"

	"github.com/GoKubes/ServerlessOrchestrator/business"
	"gorm.io/gorm"
)

type MicroservicesDAO struct {
	db *gorm.DB
}

// NewMicroservicesDAO creates a new instance of MicroservicesDAO with a database connection.
func NewMicroservicesDAO(db *gorm.DB) *MicroservicesDAO {
	return &MicroservicesDAO{db: db}
}

// GetAll retrieves all microservices from the database.
func (dao *MicroservicesDAO) GetAll() ([]business.Microservice, error) {
	var microservices []business.Microservice
	result := dao.db.Find(&microservices)
	if result.Error != nil {
		return nil, result.Error
	}
	return microservices, nil
}

// Insert adds a new microservice entity record to the database.
func (dao *MicroservicesDAO) Insert(entity interface{}) error {
	micro, ok := entity.(business.Microservice)
	if !ok {
		return errors.New("entity is not of type business.Microservice")
	}

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
func (dao *MicroservicesDAO) Delete(id uint) error {
	result := dao.db.Delete(&business.Microservice{}, id)
	return result.Error
}

// GetByID retrieves a microservice entity by its ID.
func (dao *MicroservicesDAO) GetByID(id uint) (interface{}, error) {
	var micro business.Microservice
	result := dao.db.First(&micro, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return micro, nil
}

// GetByName retrieves a microservice by its name, case-insensitively.
func (dao *MicroservicesDAO) GetByName(name string) (business.Microservice, error) {
	var micro business.Microservice
	// Using ILIKE for PostgreSQL for case-insensitive comparison
	result := dao.db.Where("LOWER(backend_name) = LOWER(?)", name).First(&micro)
	if result.Error != nil {
		return business.Microservice{}, result.Error
	}
	return micro, nil
}

// Update modifies an existing microservice record.
func (dao *MicroservicesDAO) Update(micro business.Microservice) error {
	result := dao.db.Save(&micro)
	return result.Error
}

// Ensure that MicroservicesDAO implements the MicroservicesDAO_IF interface.
var _ business.DAO_IF = (*MicroservicesDAO)(nil)
