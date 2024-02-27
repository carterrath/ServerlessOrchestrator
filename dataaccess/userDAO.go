package dataaccess

import (
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
func (userdao *UserDAO) CreateUser(user *business.User) error {
	if err := userdao.db.Create(user).Error; err != nil {
		return err
	}
	return nil
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

// Ensure that MicroservicesDAOpq implements the MicroservicesDAO_IF interface.
//var _ business.DAO_IF = &UserDAO{}
