package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"mini-douyin/model/request"
	"mini-douyin/model/response"
	"mini-douyin/service"
	"net/http"
)

type IUserController interface {
	GetUserInfo(c *gin.Context) // 获取当前登录用户信息
	UserLogin(c *gin.Context)   // 用户登录
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

// GetUserInfo 获取用户信息
func (uc UserController) GetUserInfo(c *gin.Context) {
	var userInfo request.InfoRequest
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

// UserLogin 用户登录
func (uc UserController) UserLogin(c *gin.Context) {
	var userLogin request.UserLoginRequest
	err := c.ShouldBindQuery(&userLogin)
	if err != nil {
		log.Printf("参数错误: %v", err)
		return
	}

	c.JSON(http.StatusOK, response.LoginResponse{
		Response: response.Response{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		UserId: 1,
		Token:  "ezreal",
	})
}
