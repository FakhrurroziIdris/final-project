package routers

import (
	"final-project/configs"
	"final-project/controllers"
	"final-project/helpers"
	"final-project/repositories"
	"final-project/services"

	"github.com/gin-gonic/gin"
)

func PhotosRoute(route *gin.Engine) {
	order := route.Group("/photos")

	db := helpers.PgSqlDB(configs.GetEnv().Database)
	userRepo := repositories.UserRepository(db)
	userService := services.UserService(userRepo)
	repo := repositories.PhotoRepository(db)
	service := services.PhotoService(repo)
	controller := controllers.PhotoController(service)

	order.POST("", controllers.Authentication(userService), controller.Create)
	order.GET("", controllers.Authentication(userService), controller.Get)
	order.PUT("/:photoId", controllers.Authentication(userService), controller.Update)
	order.DELETE("/:photoId", controllers.Authentication(userService), controller.Delete)
	// order.PUT("", controllers.Authentication(userService), controller.Update)
	// order.DELETE("", controllers.Authentication(userService), controller.Delete)
}
