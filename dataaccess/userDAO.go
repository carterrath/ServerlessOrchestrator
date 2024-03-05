package dataaccess

import (
	"errors"

	"github.com/GoKubes/ServerlessOrchestrator/business"
	"gorm.io/gorm"
)

// UserDAO represents the data access layer for user-related operations.
type UserDAO struct {
	db *gorm.DB
}

// NewUserDAO creates a new instance of UserDAO.
func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{db: db}
}

// CreateUser creates a new user in the database.
// was CreateUser
func (userdao *UserDAO) Insert(entity interface{}) error {
	user, ok := entity.(*business.User)
	if !ok {
		return errors.New("entity is not of type *business.User")
	}
	return userdao.db.Create(user).Error
}

// was GetUserByID
// GetByID retrieves an entity (User) by its ID.
func (userdao *UserDAO) GetByID(id uint) (interface{}, error) {
	var user business.User
	if err := userdao.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserByUsername retrieves a user from the database by username.
func (userdao *UserDAO) GetUserByUsername(username string) (*business.User, error) {
	var user business.User
	if err := userdao.db.Where("Username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserByEmail retrieves a user from the database by email.
func (userdao *UserDAO) GetUserByEmail(email string) (*business.User, error) {
	var user business.User
	if err := userdao.db.Where("Email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// CheckUsernameAndPassword checks if the provided username and password match in the database.
func (userdao *UserDAO) CheckUsernameAndPassword(username, password string) (*business.User, error) {
	var user business.User
	if err := userdao.db.Where("Username = ? AND Password = ?", username, password).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Delete removes an entity (User) record from the database by ID.
func (userdao *UserDAO) Delete(id uint) error {
	return userdao.db.Delete(&business.User{}, id).Error
}

// Ensure that MicroservicesDAO implements the MicroservicesDAO_IF interface.
var _ business.DAO_IF = (*UserDAO)(nil)
