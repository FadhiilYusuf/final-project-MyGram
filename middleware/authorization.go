package middleware

import (
	"mygram/database"
	"mygram/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Authorization(param string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param(param))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"error":   "Bad request",
				"message": "Invalid ID",
			})
			return
		}

		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userID := int(userData["id"].(float64))

		if param == "photoId" {
			db := database.GetDB()
			photo := models.Photo{}

			err := db.Select("user_id").First(&photo, id).Error

			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
					"code":    http.StatusNotFound,
					"error":   "Not Found",
					"message": "Data doesn't exist",
				})
				return
			}

			if userID != photo.UserID {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"code":    http.StatusUnauthorized,
					"error":   "Unauthorized",
					"message": "You are not allowed to access this data",
				})
				return
			}

		} else if param == "commentId" {
			db := database.GetDB()
			comment := models.Comment{}

			err := db.Select("user_id").First(&comment, id).Error

			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
					"code":    http.StatusNotFound,
					"error":   "Not Found",
					"message": "Data doesn't exist",
				})
				return
			}

			if userID != comment.UserID {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"code":    http.StatusUnauthorized,
					"error":   "Unauthorized",
					"message": "You are not allowed to access this data",
				})
				return
			}
		} else if param == "socialMediaId" {
			db := database.GetDB()
			sosmed := models.SocialMedia{}

			err := db.Select("user_id").First(&sosmed, id).Error

			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
					"code":    http.StatusNotFound,
					"error":   "Not Found",
					"message": "Data doesn't exist",
				})
				return
			}

			if userID != sosmed.UserID {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"code":    http.StatusUnauthorized,
					"error":   "Unauthorized",
					"message": "You are not allowed to access this data",
				})
				return
			}

		} else {
			if userID != id {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"code":    http.StatusUnauthorized,
					"error":   "Unauthorized",
					"message": "You are not allowed to access this data",
				})
				return
			}
		}

		ctx.Next()
	}
}
