package repositories

import "assignment-3/models"

type IRepository interface {
	UpdateSensor()
	Get() models.SensorData
}

type IDB interface {
	Get()
	Update()
}
