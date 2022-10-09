package repositories

import "final-project/models"

type IRepository interface {
	UpdateSensor()
	Get() models.SensorData
}

type IDB interface {
	Get()
	Update()
}
