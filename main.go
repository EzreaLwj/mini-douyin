package main

import (
	"fmt"
	"log"
	"mini-douyin/common"
	"mini-douyin/config"
	"mini-douyin/routes"
)

func main() {

	// 加载配置文件到全局配置结构体
	config.InitConfig()

	// 初始化数据库
	config.InitMysql()

	// 初始化配置
	common.Init()

	// 注册所有路由
	r := routes.InitRoutes()

	host := config.Conf.System.Host
	port := config.Conf.System.Port
	log.Printf("服务器在%v端口监听", port)
	r.Run(fmt.Sprintf("%s:%d", host, port))
	//srv := &http.Server{
	//	Addr:    fmt.Sprintf("%s:%d", host, port),
	//	Handler: r,
	//}
	//
	//go func() {

	//	if err := srv.ListenAndServe(); err != nil {
	//		log.Printf("listen: %s\n", err)
	//	}
	//
	//}()
	//
	//quit := make(chan os.Signal)
	//
	//signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	//<-quit
	//log.Printf("Shutting down server...")
	//
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	//if err := srv.Shutdown(ctx); err != nil {
	//	log.Printf("Server forced to shutdown: %s", err)
	//}
	//
	//log.Printf("Server exiting!")

}
