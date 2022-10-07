package main

import (
	"assignment-3/configs"
	"assignment-3/helpers"
	"assignment-3/repositories"
	"assignment-3/routers"
	"assignment-3/services"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
)

func main() {
	configs.Load()

	c := cron.New()
	c.AddFunc(configs.GetEnv().CRON.SensorInterval, func() {
		db := helpers.JsonDB(configs.GetEnv().Database.JsonFile)
		repo := repositories.SensorRepository(db)
		sensor := services.SensorService(repo)
		sensor.UpdateSensor()
	})
	c.Start()

	router := gin.Default()
	router.LoadHTMLGlob(configs.GetEnv().WebPresentation.TemplateFolder)
	routers.SensorRoute(router)
	router.Run(configs.GetEnv().WebServer.Port)
}
