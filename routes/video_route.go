package routes

import (
	"github.com/gin-gonic/gin"
	"mini-douyin/controller"
	"mini-douyin/middleware"
)

func InitVideoRoute(r *gin.RouterGroup) gin.IRoutes {
	videoController := controller.NewVideoController()
	router := r.Group("")

	{
		router.GET("/feed/", videoController.FeedVideo)
		router.POST("/publish/action/", middleware.AuthMiddleware(), videoController.PostVideo)
		router.GET("/publish/list/", middleware.AuthMiddleware(), videoController.ListVideo)
	}

	return r
}
