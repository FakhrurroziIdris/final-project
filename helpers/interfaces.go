package helpers

import "final-project/models"

type IDB interface {
	Update()
	Get() models.SensorData
}
