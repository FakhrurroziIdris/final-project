package controllers

import (
	"final-project/constants"
	"final-project/models"
	"final-project/services"
	"final-project/utils"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userService services.IService
}

func UserController(service services.IService) *userController {
	return &userController{userService: service}
}

func (cont *userController) Get(ctx *gin.Context) {
	response, err := cont.userService.Get()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (cont *userController) Create(ctx *gin.Context) {
	payload := models.User{}
	contentType := utils.GetContentType(ctx)

	if contentType == constants.AppJSON {
		if err := ctx.ShouldBindJSON(&payload); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
	} else {
		if err := ctx.ShouldBind(&payload); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
	}
	created, err := cont.userService.Create(payload)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	user := reflect.ValueOf(created).Interface().(models.User)
	ctx.JSON(http.StatusCreated, gin.H{
		"age":      user.Age,
		"email":    user.Email,
		"id":       user.ID,
		"username": user.UserName,
	})
}
