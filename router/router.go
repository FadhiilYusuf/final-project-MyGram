package router

import (
	"mygram/controller"
	"mygram/middleware"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controller.RegisterUser)
		userRouter.POST("/login", controller.LoginUser)
		userRouter.Use(middleware.Authentication(), middleware.Authorization("userId"))
		userRouter.PUT("/:userId", controller.UpdateUserData)
		userRouter.DELETE("/:userId", controller.DeleteUserAccount)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middleware.Authentication())
		photoRouter.POST("/", controller.PostPhoto)
		photoRouter.GET("/", controller.GetPhotos)
		photoRouter.Use(middleware.Authorization("photoId"))
		photoRouter.PUT("/:photoId", controller.UpdatePhoto)
		photoRouter.DELETE("/:photoId", controller.DeletePhoto)
	}
	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(middleware.Authentication())
		commentRouter.POST("/", controller.PostComment)
		commentRouter.GET("/", controller.GetComments)
		commentRouter.Use(middleware.Authorization("commentId"))
		commentRouter.PUT("/:commentId", controller.UpdateComment)
		commentRouter.DELETE("/:commentId", controller.DeleteComment)
	}
	sosmedRouter := r.Group("/socialmedias")
	{
		sosmedRouter.Use(middleware.Authentication())
		sosmedRouter.POST("/", controller.PostSocialMedia)
		sosmedRouter.GET("/", controller.GetSocialMedias)
		sosmedRouter.Use(middleware.Authorization("socialMediaId"))
		sosmedRouter.PUT("/:socialMediaId", controller.UpdateSocialMedia)
		sosmedRouter.DELETE("/:socialMediaId", controller.DeleteSocialMedia)
	}

	return r
}
