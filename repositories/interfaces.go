package repositories

type IRepository interface {
	Get() (interface{}, error)
	Create(interface{}) (interface{}, error)
}

type IDB interface {
	Get()
	Update()
}
