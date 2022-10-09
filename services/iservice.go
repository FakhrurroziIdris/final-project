package services

import "final-project/models"

type IService interface {
	Get() models.SensorData
	UpdateSensor()
}
