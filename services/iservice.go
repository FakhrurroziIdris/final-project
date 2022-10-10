package services

type IService interface {
	Get() (interface{}, error)
	Create(interface{}) (interface{}, error)
	// Login(interface{}) (interface{}, error)
	GetOne(interface{}) (interface{}, error)
}
