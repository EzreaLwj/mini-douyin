package routes

import (
	"github.com/gin-gonic/gin"
	"mini-douyin/controller"
)

func InitVideoRoute(r *gin.RouterGroup) gin.IRoutes {
	videoController := controller.NewVideoController()
	router := r.Group("")

	{
		router.GET("/feed/", videoController.FeedVideo)
		router.POST("/publish/action/", videoController.PostVideo)
		router.GET("/publish/list/", videoController.ListVideo)
	}

	return r
}
