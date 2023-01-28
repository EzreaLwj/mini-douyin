package routes

import (
	"github.com/gin-gonic/gin"
	"mini-douyin/controller"
)

// InitCommentRoutes 评论路由
func InitCommentRoutes(r *gin.RouterGroup) gin.IRoutes {
	commentController := controller.NewCommentController()
	router := r.Group("/comment")

	{
		router.GET("/list/", commentController.ListComment)
		router.POST("/action/", commentController.CommentPost)
	}

	return r
}
