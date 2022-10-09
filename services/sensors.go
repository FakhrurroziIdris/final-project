package services

import (
	"final-project/models"
	"final-project/repositories"
)

type Sensor struct {
	repo repositories.IRepository
}

func SensorService(repo repositories.IRepository) *Sensor {
	return &Sensor{repo: repo}
}

func (service *Sensor) UpdateSensor() {
	service.repo.UpdateSensor()
}

func (service *Sensor) Get() models.SensorData {
	data := service.repo.Get()
	return data
}
