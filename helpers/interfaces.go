package helpers

import "assignment-3/models"

type IDB interface {
	Update()
	Get() models.SensorData
}
