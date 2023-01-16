package routes

import (
	"github.com/gin-gonic/gin"
	"mini-douyin/controller"
)

// InitUserRoutes 注册用户路由
func InitUserRoutes(r *gin.RouterGroup) gin.IRoutes {
	userController := controller.NewUserController()
	router := r.Group("/user")

	{
		router.GET("", userController.GetUserInfo)
	}

	return r
}
