package repositories

type IRepository interface {
	Get(interface{}) (interface{}, error)
	Create(interface{}) (interface{}, error)
	GetOne(interface{}) (interface{}, error)
	Update(interface{}) (interface{}, error)
	Delete(interface{}) (interface{}, error)
}

type IDB interface {
	Get()
	Update()
}
