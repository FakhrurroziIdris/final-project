package controllers

import (
	"final-project/constants"
	"final-project/models"
	"final-project/services"
	"final-project/utils"
	"net/http"
	"reflect"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type photoController struct {
	photoService services.IService
}

func PhotoController(service services.IService) *photoController {
	return &photoController{photoService: service}
}

func bindPayloadPhoto(ctx *gin.Context) (payload models.Photo, err error) {
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

	jwtClaims := ctx.MustGet("userData").(jwt.MapClaims)
	payload.UserID = uint(jwtClaims["id"].(float64))

	return payload, err
}

func (cont *photoController) Get(ctx *gin.Context) {
	payload := ctx.MustGet("userData").(jwt.MapClaims)
	photo := models.Photo{
		UserID: uint(payload["id"].(float64)),
	}
	response, err := cont.photoService.Get(photo)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (cont *photoController) Create(ctx *gin.Context) {
	payload, err := bindPayloadPhoto(ctx)
	if err != nil {
		return
	}

	created, err := cont.photoService.Create(payload)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	photo := reflect.ValueOf(created).Interface().(models.Photo)
	ctx.JSON(http.StatusCreated, gin.H{
		"id":         photo.ID,
		"title":      photo.Title,
		"caption":    photo.Caption,
		"photo_url":  photo.PhotoUrl,
		"user_id":    photo.UserID,
		"created_at": photo.CreatedAt,
	})
}

func (cont *photoController) Update(ctx *gin.Context) {
	payload, err := bindPayloadPhoto(ctx)
	if err != nil {
		return
	}
	photoId, err := strconv.ParseUint(ctx.Param("photoId"), 0, 0)
	if err == nil {
		payload.ID = uint(photoId)
	}

	data, err := cont.photoService.Update(payload)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	photo := reflect.ValueOf(data).Interface().(models.Photo)
	ctx.JSON(http.StatusOK, gin.H{
		"id":         photo.ID,
		"title":      photo.Title,
		"caption":    photo.Caption,
		"photo_url":  photo.PhotoUrl,
		"user_id":    photo.UserID,
		"updated_at": photo.UpdatedAt,
	})
}

func (cont *photoController) Delete(ctx *gin.Context) {
	payload, err := bindPayloadPhoto(ctx)
	if err != nil {
		return
	}
	photoId, err := strconv.ParseUint(ctx.Param("photoId"), 0, 0)
	if err == nil {
		payload.ID = uint(photoId)
	}

	data, err := cont.photoService.Delete(payload)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	db := reflect.ValueOf(data).Interface().(*gorm.DB)
	if db.RowsAffected == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "No records found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Your photo has been successfully deleted",
	})
}
