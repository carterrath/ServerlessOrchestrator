package business

// Microservice struct represents a microservice with its properties.
type Microservice struct {
	ID          int    // The ID is the primary key of the microservice
	Name        string // The Name is what the user gives it, this is what is displayed to the UI
	ServiceHook ServiceHook
	BuildScript BuildScript
	PlaceHolder PlaceHolder // The GitHub link of the microservice
}
