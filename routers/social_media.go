package routers

import (
	"final-project/configs"
	"final-project/controllers"
	"final-project/helpers"
	"final-project/repositories"
	"final-project/services"

	"github.com/gin-gonic/gin"
)

func SocialMediasRoute(route *gin.Engine) {
	order := route.Group("/socialmedias")

	db := helpers.PgSqlDB(configs.GetEnv().Database)
	userRepo := repositories.UserRepository(db)
	userService := services.UserService(userRepo)

	repo := repositories.SocialMediaRepository(db)
	service := services.SocialMediaService(repo)
	controller := controllers.SocialMediaController(service)

	order.POST("", controllers.Authentication(userService), controller.Create)
	order.GET("", controllers.Authentication(userService), controller.Get)
	order.PUT("/:socialMediaId", controllers.Authentication(userService), controller.Update)
	order.DELETE("/:socialMediaId", controllers.Authentication(userService), controller.Delete)
}
