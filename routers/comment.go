package routers

import (
	"final-project/configs"
	"final-project/controllers"
	"final-project/helpers"
	"final-project/repositories"
	"final-project/services"

	"github.com/gin-gonic/gin"
)

func CommentsRoute(route *gin.Engine) {
	order := route.Group("/comments")

	db := helpers.PgSqlDB(configs.GetEnv().Database)
	userRepo := repositories.UserRepository(db)
	userService := services.UserService(userRepo)

	repo := repositories.CommentRepository(db)
	service := services.CommentService(repo)
	controller := controllers.CommentController(service)

	order.POST("", controllers.Authentication(userService), controller.Create)
	order.GET("", controllers.Authentication(userService), controller.Get)
	order.PUT("/:commentId", controllers.Authentication(userService), controller.Update)
	order.DELETE("/:commentId", controllers.Authentication(userService), controller.Delete)
}
