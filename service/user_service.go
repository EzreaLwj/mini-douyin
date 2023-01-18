package service

import (
	"log"
	req_user "mini-douyin/model/request/user"
	resp_user "mini-douyin/model/response/user"
	"mini-douyin/repository"
	"strconv"
)

type IUserService interface {
	GetUserById(request req_user.InfoRequest) (resp_user.InfoResponse, error)
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

func (u UserService) GetUserById(request req_user.InfoRequest) (resp_user.InfoResponse, error) {
	id, i2 := strconv.Atoi(request.UserId)
	infoResponse := resp_user.InfoResponse{}

	if i2 != nil {
		log.Printf("GetUserById|格式转化错误|%v,", request)
		return infoResponse, i2
	}

	userInfo, i2 := u.UserRepository.GetUserById(int64(id))
	log.Printf("用户信息为：%v", userInfo)

	infoResponse.User = userInfo

	return infoResponse, nil
}
