package services

import "assignment-3/models"

type IService interface {
	Get() models.SensorData
	UpdateSensor()
}
