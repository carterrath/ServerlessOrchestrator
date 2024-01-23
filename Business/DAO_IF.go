package business

// MicroservicesDAO interface defines the methods for interacting with microservices.
type DataAccess_IF interface {
	GetAll() ([]Microservice, error)
	GetByName(name string) (Microservice, error)
	GetByID(id int) (Microservice, error)
	Insert(micro Microservice) error
	Update(micro Microservice) error
	Delete(micro Microservice) error
	ConnectToDB() // Add any additional methods as needed
}
