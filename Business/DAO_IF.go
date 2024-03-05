package business

// DAO_IF is an interface that defines methods for interacting with microservices data.
type DAO_IF interface {
	// GetAll retrieves all microservices from the database.
	// GetAll() ([]Microservice, error)

	// Insert adds a new microservice record to the database.
	// Insert(micro Microservice) error

	// Delete removes a microservice record from the database by ID.
	// Delete(id uint) error

	// GetByID retrieves a microservice by its ID.
	// GetByID(id uint) (Microservice, error)

	// Update modifies an existing microservice record.
	// Update(micro Microservice) error

	// Insert adds a new entity record to the database.
    Insert(entity interface{}) error

    // Delete removes an entity record from the database by ID.
    Delete(id uint) error

    // GetByID retrieves an entity by its ID.
    GetByID(id uint) (interface{}, error)
}
