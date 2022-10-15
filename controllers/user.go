package controllers

import (
	"final-project/constants"
	"final-project/models"
	"final-project/services"
	"final-project/utils"
	"net/http"
	"reflect"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type userController struct {
	userService services.IService
}

func bindPayloadUser(ctx *gin.Context) (payload models.User, err error) {
	contentType := utils.GetContentType(ctx)

	if contentType == constants.AppJSON {
		if err = ctx.ShouldBindJSON(&payload); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		}
	} else {
		if err = ctx.ShouldBind(&payload); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		}
	}

	return payload, err
}

func UserController(service services.IService) *userController {
	return &userController{userService: service}
}

func (cont *userController) Get(ctx *gin.Context) {
	payload := models.User{}
	response, err := cont.userService.Get(payload)

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

func (cont *userController) Login(ctx *gin.Context) {
	payload := models.User{}
	contentType := utils.GetContentType(ctx)
	password := ""

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

	password = payload.Password
	data, err := cont.userService.GetOne(payload)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}

	user := reflect.ValueOf(data).Interface().(models.User)
	comparePass := utils.ComparePass([]byte(user.Password), []byte(password))

	if !comparePass {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}

	token := utils.GenerateToken(user.ID, user.Email)

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (cont *userController) Update(ctx *gin.Context) {
	payload, err := bindPayloadUser(ctx)
	if err != nil {
		return
	}

	data, err := cont.userService.Update(payload)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	user := reflect.ValueOf(data).Interface().(models.User)
	ctx.JSON(http.StatusOK, gin.H{
		"id":         user.ID,
		"email":      user.Email,
		"username":   user.UserName,
		"age":        user.Age,
		"updated_at": user.UpdatedAt,
	})
}

func (cont *userController) Delete(ctx *gin.Context) {
	payload := ctx.MustGet("userData").(jwt.MapClaims)
	user := models.User{
		Email: payload["email"].(string),
		GormModel: models.GormModel{
			ID: uint(payload["id"].(float64)),
		},
	}

	_, err := cont.userService.Delete(user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Your account has been successfully deleted",
	})
}
