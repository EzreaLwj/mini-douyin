package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"mini-douyin/model/request/user"
	"mini-douyin/service"
	"net/http"
)

type IUserController interface {
	GetUserInfo(c *gin.Context) // 获取当前登录用户信息
}

type UserController struct {
	UserService service.IUserService
}

// NewUserController 构造函数
func NewUserController() IUserController {
	userService := service.NewUserService()
	userController := UserController{UserService: userService}
	return userController
}

func (uc UserController) GetUserInfo(c *gin.Context) {
	var userInfo user.InfoRequest
	err := c.ShouldBindQuery(&userInfo)
	if err != nil {
		log.Printf("参数错误: %v", err)
		return
	}
	infoResponse, err := uc.UserService.GetUserById(userInfo)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, infoResponse)
}
