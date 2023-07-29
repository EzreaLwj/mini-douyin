package routes

import (
	"github.com/gin-gonic/gin"
	"mini-douyin/controller"
	"mini-douyin/middleware"
)

// InitChatRoutes 聊天路由
func InitChatRoutes(r *gin.RouterGroup) gin.IRoutes {
	chatController := controller.NewChatController()
	router := r.Group("/message")

	{
		router.GET("/chat/", middleware.AuthMiddleware(), chatController.ListChatMessage)
		router.POST("/action/", middleware.AuthMiddleware(), chatController.ChatPost)
	}

	return r
}
