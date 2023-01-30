package service

import (
	"encoding/json"
	"github.com/jinzhu/copier"
	"log"
	"mini-douyin/model/request"
	"mini-douyin/model/response"
	"mini-douyin/repository"
	"mini-douyin/utils/jwt"
	"strconv"
)

type IUserService interface {
	GetUserById(request request.InfoRequest) (response.InfoResponse, error)                 // 获取用户id
	UserRegister(userRegisterRequest request.UserRegisterRequest) response.RegisterResponse // 用户注册
	UserLogin(r *request.UserLoginRequest) response.LoginResponse                           // 用户登录
}

type UserService struct {
	UserRepository repository.IUserRepository
}

// NewUserService  构造函数
func NewUserService() IUserService {
	userRepository := repository.NewUserRepository()
	userService := UserService{UserRepository: userRepository}
	return userService
}

// UserLogin 用户登录
func (u UserService) UserLogin(r *request.UserLoginRequest) response.LoginResponse {
	userName := r.UserName
	password := r.Password

	user, err := u.UserRepository.CheckUser(userName, password)
	if err != nil {
		log.Printf("UserLogin|登录失败|%v", err)
		return response.LoginResponse{
			Response: response.Response{
				StatusCode: 1,
				StatusMsg:  "Login fail",
			},
		}
	}
	if user.UserId == 0 {
		log.Printf("UserLogin|登录失败|%v", *r)
		return response.LoginResponse{
			Response: response.Response{
				StatusCode: 1,
				StatusMsg:  "Login fail",
			},
		}
	}

	token, err := jwt.GenToken(user.UserId)
	if err != nil {
		log.Printf("UserLogin|token获取失败|")
		return response.LoginResponse{Response: response.Response{StatusCode: 1, StatusMsg: "fail"}}
	}
	log.Printf("UserLogin|用户登录|%v", user)
	return response.LoginResponse{
		Response: response.Response{
			StatusCode: 0,
			StatusMsg:  "login success",
		},
		UserId: user.UserId,
		Token:  token,
	}
}

// GetUserById 根据id获取用户信息
func (u UserService) GetUserById(request request.InfoRequest) (response.InfoResponse, error) {
	id, i2 := strconv.Atoi(request.UserId)
	infoResponse := response.InfoResponse{}

	if i2 != nil {
		log.Printf("GetUserById|格式转化错误|%v,", request)
		return infoResponse, i2
	}

	userInfo, i2 := u.UserRepository.GetUserById(int64(id))

	var useResponse response.User
	err := copier.Copy(&useResponse, &userInfo)
	if err != nil {
		log.Printf("GetUserById|类型转换错误|%v", err)
		return response.InfoResponse{}, err
	}
	infoResponse.User = useResponse
	infoResponse.StatusMsg = "success"
	infoResponse.StatusCode = 0
	infoResponse.User.IsFollow = true
	marshal, i2 := json.Marshal(infoResponse)
	log.Printf("GetUserById|用户信息为|%v", string(marshal))
	return infoResponse, nil
}

// UserRegister 用户注册
func (u UserService) UserRegister(userRegisterRequest request.UserRegisterRequest) response.RegisterResponse {
	userName := userRegisterRequest.UserName
	password := userRegisterRequest.Password

	if len(userName) > 32 || len(password) > 32 {
		log.Printf("UserRegister|用户名或密码长度错误|%v", userRegisterRequest)
		return response.RegisterResponse{
			Response: response.Response{StatusCode: 1, StatusMsg: "用户名或密码长度错误"},
		}
	}

	user, err := u.UserRepository.CreateUser(userName, password)
	if err != nil {
		log.Printf("UserRegister|数据库插入错误|%v", err)
		return response.RegisterResponse{
			Response: response.Response{StatusCode: 1, StatusMsg: "创建失败"},
		}
	}
	token, err := jwt.GenToken(user.UserId)
	if err != nil {
		log.Printf("UserRegister|token获取失败|%v", err)
		return response.RegisterResponse{}
	}
	log.Printf("UserRegister|用户注册|%v", user)
	return response.RegisterResponse{
		Response: response.Response{StatusCode: 0, StatusMsg: "创建成功"},
		UserId:   user.UserId,
		Token:    token,
	}

}
