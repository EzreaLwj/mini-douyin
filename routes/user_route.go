package routes

import (
	"github.com/gin-gonic/gin"
	"mini-douyin/controller"
	"mini-douyin/middleware"
)

// InitUserRoutes 注册用户路由
func InitUserRoutes(r *gin.RouterGroup) gin.IRoutes {
	userController := controller.NewUserController()
	router := r.Group("/user")

	{
		router.GET("/", middleware.AuthMiddleware(), userController.GetUserInfo)
		router.POST("/login/", userController.UserLogin)
		router.POST("/register/", userController.UserRegister)
	}

	return r
}
