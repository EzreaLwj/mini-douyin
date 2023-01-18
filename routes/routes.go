package routes

import (
	"github.com/gin-gonic/gin"
	"mini-douyin/config"
	"mini-douyin/middleware"
)

func InitRoutes() *gin.Engine {
	engine := gin.Default()

	// 配置全局跨域 中间件
	engine.Use(middleware.CORSMiddleware())

	// 路由分组 添加前缀
	group := engine.Group(config.Conf.System.UrlPathPrefix)

	InitUserRoutes(group)

	return engine
}
