package controller

import (
	"mygram/database"
	"mygram/helpers"
	"mygram/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func PostSocialMedia(ctx *gin.Context) {
	db := database.GetDB()

	SocialMedia := models.SocialMedia{}
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := int(userData["id"].(float64))

	reqHeaders := helpers.GetRequestHeaders(ctx)
	if reqHeaders.ContentType == "application/json" {
		ctx.ShouldBindJSON(&SocialMedia)
	} else {
		ctx.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserID = userID

	err := db.Debug().Create(&SocialMedia).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id":               SocialMedia.ID,
		"name":             SocialMedia.Name,
		"social_media_url": SocialMedia.SocialMediaUrl,
		"user_id":          SocialMedia.UserID,
		"created_at":       SocialMedia.CreatedAt,
	})
}

func GetSocialMedias(ctx *gin.Context) {
	db := database.GetDB()

	socialMedias := []models.SocialMedia{}
	err := db.Preload(clause.Associations).Find(&socialMedias).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
		return
	}
	socialMediasResponse := []models.GetAllSosmedResponse{}

	for _, socialMedia := range socialMedias {
		response := models.GetAllSosmedResponse{}

		response.GormModel = socialMedia.GormModel
		response.Name = socialMedia.Name
		response.SocialMediaUrl = socialMedia.SocialMediaUrl
		response.UserID = socialMedia.UserID
		response.User.UserName = socialMedia.User.UserName
		response.User.Email = socialMedia.User.Email

		socialMediasResponse = append(socialMediasResponse, response)
	}

	ctx.JSON(http.StatusOK, socialMediasResponse)
}

func UpdateSocialMedia(ctx *gin.Context) {
	db := database.GetDB()

	SocialMedia := models.SocialMedia{}
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := int(userData["id"].(float64))
	SocialMediaId, err := strconv.Atoi(ctx.Param("socialMediaId"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"error":   "Bad request",
			"message": "Invalid ID",
		})
		return
	}

	reqHeaders := helpers.GetRequestHeaders(ctx)
	if reqHeaders.ContentType == "application/json" {
		ctx.ShouldBindJSON(&SocialMedia)
	} else {
		ctx.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserID = userID

	err = db.Model(&SocialMedia).Where("id=?", SocialMediaId).Updates(SocialMedia).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":               SocialMediaId,
		"name":             SocialMedia.Name,
		"social_media_url": SocialMedia.SocialMediaUrl,
		"user_id":          SocialMedia.UserID,
		"updated_at":       SocialMedia.UpdatedAt,
	})
}

func DeleteSocialMedia(ctx *gin.Context) {
	db := database.GetDB()

	SocialMedia := models.SocialMedia{}
	SocialMediaId, err := strconv.Atoi(ctx.Param("socialMediaId"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"error":   "Bad request",
			"message": "Invalid ID",
		})
		return
	}

	err = db.Where("id=?", SocialMediaId).Delete(&SocialMedia).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"error":   "Not Found",
			"message": "SocialMedia not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Your Social Media has been succsessfuly deleted",
	})

}
