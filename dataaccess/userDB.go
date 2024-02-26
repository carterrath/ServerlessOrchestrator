package dataaccess

import (
	"github.com/GoKubes/ServerlessOrchestrator/business"
	"gorm.io/gorm"
)

// UserDB represents the data access layer for user-related operations.
type UserDB struct {
	db *gorm.DB
}

// NewUserDB creates a new instance of UserDB.
func NewUserDB(db *gorm.DB) *UserDB {
	return &UserDB{db: db}
}

// AutoMigrate creates or updates the user table schema in the database.
func (udb *UserDB) AutoMigrate() error {
	// AutoMigrate creates or updates the user table schema based on the User struct.
	if err := udb.db.AutoMigrate(&business.User{}); err != nil {
		return err
	}
	return nil
}

// CreateUser creates a new user in the database.
func (udb *UserDB) CreateUser(user *business.User) error {
	if err := udb.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

// GetUserByUsername retrieves a user from the database by username.
func (udb *UserDB) GetUserByUsername(username string) (*business.User, error) {
	var user business.User
	if err := udb.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserByEmail retrieves a user from the database by email.
func (udb *UserDB) GetUserByEmail(email string) (*business.User, error) {
	var user business.User
	if err := udb.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// CheckUsernameAndPassword checks if the provided username and password match in the database.
func (udb *UserDB) CheckUsernameAndPassword(username, password string) (*business.User, error) {
	var user business.User
	if err := udb.db.Where("username = ? AND password = ?", username, password).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
