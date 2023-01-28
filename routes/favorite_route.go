package routes

import (
	"github.com/gin-gonic/gin"
	"mini-douyin/controller"
)

// InitFavoriteRoutes 注册用户路由
func InitFavoriteRoutes(r *gin.RouterGroup) gin.IRoutes {
	favoriteController := controller.NewFavoriteController()
	router := r.Group("/favorite")

	{
		router.GET("/list/", favoriteController.FavoriteVideoList)
		router.POST("/action/", favoriteController.FavoriteAction)
	}

	return r
}
