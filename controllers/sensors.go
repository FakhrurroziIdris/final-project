package controllers

import (
	"assignment-3/configs"
	"assignment-3/services"
	"net/http"

	"assignment-3/models"

	"github.com/gin-gonic/gin"
)

type sensorController struct {
	sensorService services.IService
}

func NewSensorController(service services.Sensor) *sensorController {
	return &sensorController{sensorService: &service}
}

func (cont *sensorController) Get(ctx *gin.Context) {
	var fileData = cont.sensorService.Get()

	var data = models.Status{
		Water: fileData.Water,
		Wind:  fileData.Wind,
	}

	switch {
	case data.Water <= 5:
		data.WaterLevel = "Aman"
	case data.Water >= 6 && data.Water <= 8:
		data.WaterLevel = "Siaga"
	case data.Water > 8:
		data.WaterLevel = "Bahaya"
	}

	switch {
	case data.Wind <= 6:
		data.WindLevel = "Aman"
	case data.Wind >= 7 && data.Wind <= 15:
		data.WindLevel = "Siaga"
	case data.Wind > 15:
		data.WindLevel = "Bahaya"
	}

	ctx.HTML(http.StatusOK, "template.html", gin.H{
		"Water":       data.Water,
		"WaterLevel":  data.WaterLevel,
		"Wind":        data.Wind,
		"WindLevel":   data.WindLevel,
		"PageRefresh": configs.GetEnv().WebPresentation.PageRefresh,
	})
}
