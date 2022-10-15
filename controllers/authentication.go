package controllers

import (
	"final-project/models"
	"final-project/services"
	"final-project/utils"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Authentication(service *services.User) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		verifyToken, err := utils.VerifyToken(ctx)
		_ = verifyToken

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthenticated",
				"message": err.Error(),
			})
			return
		}

		token := verifyToken.(jwt.MapClaims)
		user := models.User{
			Email: token["email"].(string),
			GormModel: models.GormModel{
				ID: uint(token["id"].(float64)),
			},
		}
		if _, err := service.GetOne(user); err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthenticated",
				"message": "user " + err.Error(),
			})
			return
		}

		ctx.Set("userData", verifyToken)
		ctx.Next()
	}
}
