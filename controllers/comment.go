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

type commentController struct {
	commentService services.IService
}

func CommentController(service services.IService) *commentController {
	return &commentController{commentService: service}
}

func bindPayloadComment(ctx *gin.Context) (payload models.Comment, err error) {
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

func (cont *commentController) Get(ctx *gin.Context) {
	payload := ctx.MustGet("userData").(jwt.MapClaims)
	comment := models.Comment{
		UserID: uint(payload["id"].(float64)),
	}
	response, err := cont.commentService.Get(comment)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	comments := []models.Comment{}
	photo := models.Photo{}
	user := models.User{}
	for k, v := range response.(map[string]interface{}) {
		switch k {
		case "comments":
			comments = v.([]models.Comment)
			break
		case "photo":
			photo = v.(models.Photo)
			break
		case "user":
			user = v.(models.User)
			break
		}
	}

	models := []gin.H{}
	for i := 0; i < len(comments); i++ {
		models = append(models, gin.H{
			"id":         comments[i].ID,
			"message":    comments[i].Message,
			"photo_id":   comments[i].PhotoID,
			"user_id":    comments[i].UserID,
			"updated_at": comments[i].UpdatedAt,
			"created_at": comments[i].CreatedAt,
			"User": map[string]interface{}{
				"id":       user.ID,
				"email":    user.Email,
				"username": user.UserName,
			},
			"Photo": map[string]interface{}{
				"id":        photo.ID,
				"title":     photo.Title,
				"caption":   photo.Caption,
				"photo_url": photo.PhotoUrl,
				"user_id":   photo.UserID,
			},
		})
	}
	ctx.JSON(http.StatusOK, models)
}

func (cont *commentController) Create(ctx *gin.Context) {
	payload, err := bindPayloadComment(ctx)
	if err != nil {
		return
	}

	created, err := cont.commentService.Create(payload)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	comment := reflect.ValueOf(created).Interface().(models.Comment)
	ctx.JSON(http.StatusCreated, gin.H{
		"id":         comment.ID,
		"message":    comment.Message,
		"photo_id":   comment.PhotoID,
		"user_id":    comment.UserID,
		"created_at": comment.CreatedAt,
	})
}

func (cont *commentController) Update(ctx *gin.Context) {
	payload, err := bindPayloadComment(ctx)
	if err != nil {
		return
	}
	commentId, err := strconv.ParseUint(ctx.Param("commentId"), 0, 0)
	if err == nil {
		payload.ID = uint(commentId)
	}

	data, err := cont.commentService.Update(payload)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	comment := models.Comment{}
	photo := models.Photo{}
	for k, v := range data.(map[string]interface{}) {
		switch k {
		case "comments":
			comment = v.(models.Comment)
			break
		case "photo":
			photo = v.(models.Photo)
			break
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":         comment.ID,
		"title":      photo.Title,
		"caption":    photo.Caption,
		"photo_url":  photo.PhotoUrl,
		"user_id":    comment.UserID,
		"updated_at": comment.UpdatedAt,
	})
}

func (cont *commentController) Delete(ctx *gin.Context) {
	payload, err := bindPayloadComment(ctx)
	if err != nil {
		return
	}
	commentId, err := strconv.ParseUint(ctx.Param("commentId"), 0, 0)
	if err == nil {
		payload.ID = uint(commentId)
	}

	data, err := cont.commentService.Delete(payload)
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
		"message": "Your comment has been successfully deleted",
	})
}
