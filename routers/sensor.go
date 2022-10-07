package routers

import (
	"assignment-3/configs"
	"assignment-3/controllers"
	"assignment-3/helpers"
	"assignment-3/repositories"
	"assignment-3/services"

	"github.com/gin-gonic/gin"
)

func authentication() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	})
}

func SensorRoute(route *gin.Engine) {
	order := route.Group("/sensor")

	db := helpers.JsonDB(configs.GetEnv().Database.JsonFile)
	repo := repositories.SensorRepository(db)
	service := services.SensorService(repo)
	sensorController := controllers.NewSensorController(*service)
	order.GET("", sensorController.Get)
}
