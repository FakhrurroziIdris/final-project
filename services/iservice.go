package services

type IService interface {
	Get() (interface{}, error)
	Create(interface{}) (interface{}, error)
}
