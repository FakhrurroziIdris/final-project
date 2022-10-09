package repositories

import (
	"final-project/helpers"
	"final-project/models"
)

type Sensor struct {
	db helpers.IDB
}

func SensorRepository(db helpers.IDB) *Sensor {
	return &Sensor{db: db}
}

func (repo *Sensor) UpdateSensor() {
	repo.db.Update()
}

func (repo *Sensor) Get() models.SensorData {
	return repo.db.Get()
}
