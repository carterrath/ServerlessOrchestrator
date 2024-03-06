package business

// DAO_IF is an interface that defines methods for interacting with microservices data.
type DAO_IF interface {
	// Insert adds a new entity record to the database.
	Insert(entity interface{}) error

	// Delete removes an entity record from the database by ID.
	Delete(id uint) error

	// GetByID retrieves an entity by its ID.
	GetByID(id uint) (interface{}, error)
}
