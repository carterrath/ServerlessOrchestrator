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

type MicroserviceData struct {
	ID            uint
	CreatedAt     string
	UpdatedAt     string
	DeletedAt     string
	FriendlyName  string
	RepoLink      string
	StatusMessage string
	IsActive      bool
	User          UserData
	Inputs        []business.Input
	OutputLink    string
	BackendName   string `gorm:"unique"`
	ImageID       string
}

type UserData struct {
	ID        uint
	CreatedAt string
	UpdatedAt string
	DeletedAt string
	Username  string
	Email     string
}

// GetAllWithUsers retrieves all microservices with their associated users.
func (dao *MicroservicesDAO) GetAllWithUsers(userDAO *UserDAO) ([]MicroserviceData, error) {
	microservices, err := dao.GetAll()
	if err != nil {
		return nil, err
	}

	// Batch fetch all user IDs
	userIDs := make([]uint, len(microservices))
	for i, microservice := range microservices {
		userIDs[i] = microservice.UserID
	}

	users, err := userDAO.GetBatchByID(userIDs)
	if err != nil {
		return nil, err
	}

	// Map user IDs to users for quick lookup
	userMap := make(map[uint]business.User)
	for _, user := range users {
		userMap[user.ID] = user
	}

	var microservicesData []MicroserviceData
	for _, microservice := range microservices {
		user, ok := userMap[microservice.UserID]
		if !ok {
			continue
		}

		microservicesData = append(microservicesData, MicroserviceData{
			ID:            microservice.ID,
			CreatedAt:     microservice.CreatedAt.String(),
			UpdatedAt:     microservice.UpdatedAt.String(),
			DeletedAt:     microservice.DeletedAt.Time.String(),
			FriendlyName:  microservice.FriendlyName,
			RepoLink:      microservice.RepoLink,
			StatusMessage: microservice.StatusMessage,
			IsActive:      microservice.IsActive,
			User:          UserData{ID: user.ID, CreatedAt: user.CreatedAt.String(), UpdatedAt: user.UpdatedAt.String(), DeletedAt: user.DeletedAt.Time.String(), Username: user.Username, Email: user.Email},
			Inputs:        microservice.Inputs,
			OutputLink:    microservice.OutputLink,
			BackendName:   microservice.BackendName,
			ImageID:       microservice.ImageID,
		})
	}

	return microservicesData, nil
}

func (dao *MicroservicesDAO) UpdateStatusMessage(id uint, message string) error {
	result := dao.db.Model(&business.Microservice{}).Where("ID = ?", id).Update("status_message", message)
	return result.Error
}

func (dao *MicroservicesDAO) GetAllActiveCount() (int, error) {
	var count int64
	result := dao.db.Model(&business.Microservice{}).Where("is_active = ?", true).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(count), nil // Convert int64 to int before returning
}

func (dao *MicroservicesDAO) GetStatusMessage(id uint) (string, error) {
	var micro business.Microservice
	result := dao.db.Select("status_message").First(&micro, id)
	return micro.StatusMessage, result.Error
}

// Ensure that MicroservicesDAO implements the MicroservicesDAO_IF interface.
var _ business.DAO_IF = (*MicroservicesDAO)(nil)
