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

type socialMediaController struct {
	socialMediaService services.IService
}

func SocialMediaController(service services.IService) *socialMediaController {
	return &socialMediaController{socialMediaService: service}
}

func bindPayloadSocialMedia(ctx *gin.Context) (payload models.SocialMedia, err error) {
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

func (cont *socialMediaController) Get(ctx *gin.Context) {
	payload := ctx.MustGet("userData").(jwt.MapClaims)
	socialMedia := models.SocialMedia{
		UserID: uint(payload["id"].(float64)),
	}
	response, err := cont.socialMediaService.Get(socialMedia)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	social_medias := []models.SocialMedia{}
	user := models.User{}

	for k, v := range response.(map[string]interface{}) {
		switch k {
		case "social_medias":
			social_medias = v.([]models.SocialMedia)
			break
		case "user":
			user = v.(models.User)
			break
		}
	}

	data := []map[string]interface{}{}
	for i := 0; i < len(social_medias); i++ {
		data = append(data, map[string]interface{}{
			"id":               social_medias[i].ID,
			"name":             social_medias[i].Name,
			"social_media_url": social_medias[i].SocialMediaUrl,
			"UserId":           social_medias[i].UserID,
			"createdAt":        social_medias[i].CreatedAt,
			"updatedAt":        social_medias[i].UpdatedAt,
			"User": map[string]interface{}{
				"id":                user.ID,
				"username":          user.UserName,
				"profile_image_url": user.Photos[0].PhotoUrl,
			},
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"social_medias": data,
	})
}

func (cont *socialMediaController) Create(ctx *gin.Context) {
	payload, err := bindPayloadSocialMedia(ctx)
	if err != nil {
		return
	}

	created, err := cont.socialMediaService.Create(payload)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	socialMedia := reflect.ValueOf(created).Interface().(models.SocialMedia)
	ctx.JSON(http.StatusCreated, gin.H{
		"id":               socialMedia.ID,
		"name":             socialMedia.Name,
		"social_media_url": socialMedia.SocialMediaUrl,
		"user_id":          socialMedia.UserID,
		"created_at":       socialMedia.CreatedAt,
	})
}

func (cont *socialMediaController) Update(ctx *gin.Context) {
	payload, err := bindPayloadSocialMedia(ctx)
	if err != nil {
		return
	}
	socialMediaId, err := strconv.ParseUint(ctx.Param("socialMediaId"), 0, 0)
	if err == nil {
		payload.ID = uint(socialMediaId)
	}

	data, err := cont.socialMediaService.Update(payload)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	socialMedia := reflect.ValueOf(data).Interface().(models.SocialMedia)
	ctx.JSON(http.StatusOK, gin.H{
		"id":               socialMedia.ID,
		"name":             socialMedia.Name,
		"social_media_url": socialMedia.SocialMediaUrl,
		"user_id":          socialMedia.UserID,
		"updated_at":       socialMedia.UpdatedAt,
	})
}

func (cont *socialMediaController) Delete(ctx *gin.Context) {
	payload, err := bindPayloadSocialMedia(ctx)
	if err != nil {
		return
	}
	socialMediaId, err := strconv.ParseUint(ctx.Param("socialMediaId"), 0, 0)
	if err == nil {
		payload.ID = uint(socialMediaId)
	}

	data, err := cont.socialMediaService.Delete(payload)
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
		"message": "Your social media has been successfully deleted",
	})
}
