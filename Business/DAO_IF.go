package Business

type Microservice struct {
	ID          int    // The ID is the primary key of the microservice
	Name        string // The Name is what the user gives it, this is what is displayed to the UI
	ServiceHook string
	BuildScript string
	PlaceHolder string // The GitHub link of the microservice
}

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
