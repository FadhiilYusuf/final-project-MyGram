package controller

import (
	"fmt"
	"mygram/database"
	"mygram/helpers"
	"mygram/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func PostComment(ctx *gin.Context) {
	db := database.GetDB()

	Comment := models.Comment{}
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := int(userData["id"].(float64))

	reqHeaders := helpers.GetRequestHeaders(ctx)
	if reqHeaders.ContentType == "application/json" {
		ctx.ShouldBindJSON(&Comment)
	} else {
		ctx.ShouldBind(&Comment)
	}

	Comment.UserID = userID

	err := db.Debug().Create(&Comment).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id":         Comment.ID,
		"message":    Comment.Message,
		"photo_id":   Comment.PhotoID,
		"user_id":    Comment.UserID,
		"created_at": Comment.CreatedAt,
	})
}

func GetComments(ctx *gin.Context) {
	fmt.Println("ini diatas")
	db := database.GetDB()
	fmt.Println("ini dibawah")
	comments := []models.Comment{}
	err := db.Preload(clause.Associations).Find(&comments).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
		return
	}
	commentsResponse := []models.GetAllCommentsResponse{}

	for _, comment := range comments {
		response := models.GetAllCommentsResponse{}

		response.GormModel = comment.GormModel
		response.Message = comment.Message
		response.PhotoID = comment.PhotoID
		response.UserID = comment.UserID
		response.Photo.ID = comment.Photo.ID
		response.Photo.Title = comment.Photo.Title
		response.Photo.Caption = comment.Photo.Caption
		response.Photo.PhotoUrl = comment.Photo.PhotoUrl
		response.User.UserName = comment.User.UserName
		response.User.Email = comment.User.Email

		commentsResponse = append(commentsResponse, response)
	}

	ctx.JSON(http.StatusOK, commentsResponse)
}

func UpdateComment(ctx *gin.Context) {
	db := database.GetDB()

	Comment := models.Comment{}
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := int(userData["id"].(float64))
	CommentId, err := strconv.Atoi(ctx.Param("commentId"))
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
		ctx.ShouldBindJSON(&Comment)
	} else {
		ctx.ShouldBind(&Comment)
	}

	Comment.UserID = userID

	err = db.Model(&Comment).Where("id=?", CommentId).Updates(Comment).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":         CommentId,
		"message":    Comment.Message,
		"photo_id":   Comment.PhotoID,
		"user_id":    Comment.UserID,
		"updated_at": Comment.UpdatedAt,
	})
}

func DeleteComment(ctx *gin.Context) {
	db := database.GetDB()

	Comment := models.Comment{}
	CommentId, err := strconv.Atoi(ctx.Param("commentId"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"error":   "Bad request",
			"message": "Invalid ID",
		})
		return
	}

	err = db.Where("id=?", CommentId).Delete(&Comment).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"error":   "Not Found",
			"message": "Comment not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Your Comment has been succsessfuly deleted",
	})

}
