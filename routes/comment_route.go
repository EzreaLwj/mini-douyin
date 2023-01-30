package routes

import (
	"github.com/gin-gonic/gin"
	"mini-douyin/controller"
	"mini-douyin/middleware"
)

// InitCommentRoutes 评论路由
func InitCommentRoutes(r *gin.RouterGroup) gin.IRoutes {
	commentController := controller.NewCommentController()
	router := r.Group("/comment")

	{
		router.GET("/list/", middleware.AuthMiddleware(), commentController.ListComment)
		router.POST("/action/", middleware.AuthMiddleware(), commentController.CommentPost)
	}

	return r
}
