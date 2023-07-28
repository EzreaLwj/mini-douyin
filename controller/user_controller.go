package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"mini-douyin/model/request"
	"mini-douyin/service"
	"net/http"
)

type IUserController interface {
	GetUserInfo(c *gin.Context)  // 获取当前登录用户信息
	UserLogin(c *gin.Context)    // 用户登录
	UserRegister(c *gin.Context) // 用户注册
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
	loginResponse := uc.UserService.UserLogin(&userLogin)
	c.JSON(http.StatusOK, loginResponse)
}

// UserRegister 用户注册
func (uc UserController) UserRegister(c *gin.Context) {
	var userRegister request.UserRegisterRequest
	err := c.ShouldBindQuery(&userRegister)
	if err != nil {
		log.Printf("参数错误: %v", err)
		return
	}
	registerResponse := uc.UserService.UserRegister(userRegister)
	c.JSON(http.StatusOK, registerResponse)
}
