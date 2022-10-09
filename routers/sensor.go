package routers

import (
	"final-project/configs"
	"final-project/controllers"
	"final-project/helpers"
	"final-project/repositories"
	"final-project/services"

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
