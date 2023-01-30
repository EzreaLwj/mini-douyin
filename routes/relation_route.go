package routes

import (
	"github.com/gin-gonic/gin"
	"mini-douyin/controller"
	"mini-douyin/middleware"
)

// InitRelationRoutes 关注用户路由
func InitRelationRoutes(r *gin.RouterGroup) gin.IRoutes {
	relationController := controller.NewRelationController()
	router := r.Group("/relation")

	{
		router.GET("/follower/list/", middleware.AuthMiddleware(), relationController.GetFollower)
		router.GET("/follow/list/", middleware.AuthMiddleware(), relationController.GetFollow)
		router.POST("/action/", middleware.AuthMiddleware(), relationController.FollowAction)
	}

	return r
}
