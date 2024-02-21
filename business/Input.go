package business

type Input struct {
	ID             uint
	MicroserviceID uint // This field represents the foreign key
	Name           string
	DataType       string
}
