package routes

import (
	"github.com/gin-gonic/gin"
	"mini-douyin/controller"
)

// InitRelationRoutes 关注用户路由
func InitRelationRoutes(r *gin.RouterGroup) gin.IRoutes {
	relationController := controller.NewRelationController()
	router := r.Group("/relation")

	{
		router.GET("/follower/list/", relationController.GetFollower)
		router.GET("/follow/list/", relationController.GetFollow)
		router.POST("/action/", relationController.FollowAction)
	}

	return r
}
