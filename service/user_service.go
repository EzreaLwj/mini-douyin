package service

import (
	"encoding/json"
	"github.com/jinzhu/copier"
	"log"
	"mini-douyin/model/request"
	"mini-douyin/model/response"
	"mini-douyin/repository"
	"strconv"
)

type IUserService interface {
	GetUserById(request request.InfoRequest) (response.InfoResponse, error)
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
