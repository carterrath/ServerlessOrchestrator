package Business

// MicroservicesDAO interface defines the methods for interacting with microservices.
type DataAccess_IF interface {
	GetAllMicroservices() ([]Microservice, error)
	GetMicroserviceByName(name string) (Microservice, error)
	GetMicroserviceByID(id int) (Microservice, error)
	InsertMicroservice(micro Microservice) error
	UpdateMicroservice(micro Microservice) error
	DeleteMicroservice(micro Microservice) error
	ConnectToDB() // Add any additional methods as needed
}
