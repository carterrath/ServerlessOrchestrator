package business

// Microservice struct represents a microservice with its properties.
type Microservice struct {
	FriendlyName string
	RepoLink     string
	Status       string
	UserID       uint
	Inputs       []Input
	OutputLink   string
	BackendName  string 
}
