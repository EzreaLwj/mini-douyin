package routes

import (
	"github.com/gin-gonic/gin"
	"mini-douyin/controller"
	"mini-douyin/middleware"
)

// InitFavoriteRoutes 注册用户路由
func InitFavoriteRoutes(r *gin.RouterGroup) gin.IRoutes {
	favoriteController := controller.NewFavoriteController()
	router := r.Group("/favorite")

	{
		router.GET("/list/", middleware.AuthMiddleware(), favoriteController.FavoriteVideoList)
		router.POST("/action/", middleware.AuthMiddleware(), favoriteController.FavoriteAction)
	}

	return r
}
