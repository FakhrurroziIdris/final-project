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

func UsersRoute(route *gin.Engine) {
	order := route.Group("/users")

	db := helpers.PgSqlDB(configs.GetEnv().Database)
	repo := repositories.UserRepository(db)
	service := services.UserService(repo)
	controller := controllers.UserController(service)

	// order.GET("", utils.Authentication(), controller.Get)
	order.POST("/register", controller.Create)
	order.POST("/login", controller.Login)
	order.PUT("", controllers.Authentication(service), controller.Update)
	order.DELETE("", controllers.Authentication(service), controller.Delete)
}
