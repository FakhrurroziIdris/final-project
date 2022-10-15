package main

import (
	"final-project/configs"
	"final-project/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	configs.Load()

	router := gin.Default()
	routers.UsersRoute(router)
	routers.PhotosRoute(router)
	routers.CommentsRoute(router)
	router.Run(configs.GetEnv().WebServer.Port)
}
